package main

import (
	_ "embed"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	poll "github.com/unickorn/golem-poll/poll_out"
	"time"
)

// env is not working ... :(
// TODO: unhardcode this
var secret = "secret-code!"

func init() {
	i := GolemPoll{}
	poll.SetExportsIlhanGolemPollApi(i)
}

var data poll.IlhanGolemPollApiPollData

// GolemPoll is the implementation of the poll interface, a poll template for Golem Cloud.
type GolemPoll struct {
}

// Create handles a base64 encoded poll creation request.
func (g GolemPoll) Create(p poll.IlhanGolemPollApiPoll) poll.Result[poll.IlhanGolemPollApiPoll, string] {
	if len(data.Poll.Questions) > 0 {
		return poll.Result[poll.IlhanGolemPollApiPoll, string]{
			Kind: poll.Err,
			Err:  "Failed to create poll, this poll title is taken!",
		}
	}
	if len(p.Questions) == 0 {
		return poll.Result[poll.IlhanGolemPollApiPoll, string]{
			Kind: poll.Err,
			Err:  "Failed to create poll, questions cannot be empty.",
		}
	}
	endTime := time.Unix(p.End, 0)
	if endTime.Before(time.Now()) {
		return poll.Result[poll.IlhanGolemPollApiPoll, string]{
			Kind: poll.Err,
			Err:  "Failed to create poll, end time cannot be in the past",
		}
	}
	data = poll.IlhanGolemPollApiPollData{
		Poll:      p,
		Open:      true,
		Responses: make([]poll.IlhanGolemPollApiResponse, 0),
	}
	return poll.Result[poll.IlhanGolemPollApiPoll, string]{
		Kind: poll.Ok,
		Val:  data.Poll,
	}
}

// Respond responds to the poll.
func (g GolemPoll) Respond(response poll.IlhanGolemPollApiResponseCreation) poll.Result[struct{}, string] {
	if len(data.Poll.Questions) == 0 {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "Poll not found.",
		}
	}
	accessToken := response.UserToken
	email, ok := checkAccess(accessToken)
	if !ok {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "Unauthorized",
		}
	}
	answers := response.Answers
	if len(data.Poll.Questions) != len(answers) {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "Answers length does not match questions length.",
		}
	}
	if !data.Open {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "This poll has been closed.",
		}
	}
	endTime := time.Unix(data.Poll.End, 0)
	if endTime.Before(time.Now()) {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "This poll has ended.",
		}
	}
	for k, r := range data.Responses {
		if r.User == email {
			data.Responses[k].Answers = answers
			return poll.Result[struct{}, string]{
				Kind: poll.Ok,
			}
		}
	}
	data.Responses = append(data.Responses, poll.IlhanGolemPollApiResponse{
		User:    email,
		Answers: answers,
	})
	return poll.Result[struct{}, string]{
		Kind: poll.Ok,
	}
}

// Get returns the poll results and the poll itself.
func (g GolemPoll) Get(accessToken string) poll.Result[poll.IlhanGolemPollApiPollData, string] {
	if len(data.Poll.Questions) == 0 {
		return poll.Result[poll.IlhanGolemPollApiPollData, string]{
			Kind: poll.Err,
			Err:  "Poll not found.",
		}
	}
	email, ok := checkAccess(accessToken)
	if !ok {
		return poll.Result[poll.IlhanGolemPollApiPollData, string]{
			Kind: poll.Err,
			Err:  "Unauthorized",
		}
	}
	if data.Poll.OwnerEmail != email {
		return poll.Result[poll.IlhanGolemPollApiPollData, string]{
			Kind: poll.Err,
			Err:  "Only the owner can get the poll results!",
		}
	}
	return poll.Result[poll.IlhanGolemPollApiPollData, string]{
		Kind: poll.Ok,
		Val:  data,
	}
}

// Close closes the poll. No more responses can be made.
func (g GolemPoll) Close(accessToken string) poll.Result[struct{}, string] {
	if len(data.Poll.Questions) == 0 {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "Poll not found.",
		}
	}
	email, ok := checkAccess(accessToken)
	if !ok {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "Unauthorized",
		}
	}
	if data.Poll.OwnerEmail != email {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "Only the owner can close the poll!",
		}
	}
	if !data.Open {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "This poll has already been closed.",
		}
	}
	data.Open = false
	return poll.Result[struct{}, string]{
		Kind: poll.Ok,
	}
}

// Delete deletes the poll and the responses. This allows the form to be re-created with the same ID.
func (g GolemPoll) Delete(accessToken string) poll.Result[struct{}, string] {
	if len(data.Poll.Questions) == 0 {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "Poll not found.",
		}
	}
	email, ok := checkAccess(accessToken)
	if !ok {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "Unauthorized",
		}
	}
	if data.Poll.OwnerEmail != email {
		return poll.Result[struct{}, string]{
			Kind: poll.Err,
			Err:  "Only the owner can delete the poll!",
		}
	}
	data = poll.IlhanGolemPollApiPollData{}
	return poll.Result[struct{}, string]{
		Kind: poll.Ok,
	}
}

// checkAccess validates an access token with the public key and returns the user's email if valid. If not, the second value
// will be false.
func checkAccess(accessToken string) (string, bool) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil || !token.Valid {
		return "", false
	}
	return token.Claims.(jwt.MapClaims)["sub"].(string), token.Valid
}

func main() {
}
