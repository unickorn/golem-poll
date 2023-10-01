// Generated by `wit-bindgen` 0.11.0. DO NOT EDIT!
package poll

// #include "poll.h"
import "C"

import "unsafe"

type WasiCliEnvironmentTuple2StringStringT struct {
  F0 string
  F1 string
}

type IlhanGolemPollApiResponseCreation struct {
  UserToken string
  Answers []uint32
}

type IlhanGolemPollApiQuestion struct {
  Text string
  Answers []string
}

type IlhanGolemPollApiPoll struct {
  OwnerEmail string
  Title string
  End int64
  Questions []IlhanGolemPollApiQuestion
}

type IlhanGolemPollApiResponse struct {
  User string
  Answers []uint32
}

type IlhanGolemPollApiPollData struct {
  Poll IlhanGolemPollApiPoll
  Open bool
  Responses []IlhanGolemPollApiResponse
}

// Import functions from wasi:cli/environment
func WasiCliEnvironmentGetEnvironment() []WasiCliEnvironmentTuple2StringStringT {
  var ret C.poll_list_tuple2_string_string_t
  C.wasi_cli_environment_get_environment(&ret)
  var lift_ret []WasiCliEnvironmentTuple2StringStringT
  lift_ret = make([]WasiCliEnvironmentTuple2StringStringT, ret.len)
  if ret.len > 0 {
    for lift_ret_i := 0; lift_ret_i < int(ret.len); lift_ret_i++ {
      var empty_lift_ret C.poll_tuple2_string_string_t
      lift_ret_ptr := *(*C.poll_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ret.ptr)) +
      uintptr(lift_ret_i)*unsafe.Sizeof(empty_lift_ret)))
      var list_lift_ret WasiCliEnvironmentTuple2StringStringT
      var list_lift_ret_F0 string
      list_lift_ret_F0 = C.GoStringN(lift_ret_ptr.f0.ptr, C.int(lift_ret_ptr.f0.len))
      list_lift_ret.F0 = list_lift_ret_F0
      var list_lift_ret_F1 string
      list_lift_ret_F1 = C.GoStringN(lift_ret_ptr.f1.ptr, C.int(lift_ret_ptr.f1.len))
      list_lift_ret.F1 = list_lift_ret_F1
      lift_ret[lift_ret_i] = list_lift_ret
    }
  }
  return lift_ret
}

func WasiCliEnvironmentGetArguments() []string {
  var ret C.poll_list_string_t
  C.wasi_cli_environment_get_arguments(&ret)
  var lift_ret []string
  lift_ret = make([]string, ret.len)
  if ret.len > 0 {
    for lift_ret_i := 0; lift_ret_i < int(ret.len); lift_ret_i++ {
      var empty_lift_ret C.poll_string_t
      lift_ret_ptr := *(*C.poll_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ret.ptr)) +
      uintptr(lift_ret_i)*unsafe.Sizeof(empty_lift_ret)))
      var list_lift_ret string
      list_lift_ret = C.GoStringN(lift_ret_ptr.ptr, C.int(lift_ret_ptr.len))
      lift_ret[lift_ret_i] = list_lift_ret
    }
  }
  return lift_ret
}

func WasiCliEnvironmentInitialCwd() Option[string] {
  var ret C.poll_option_string_t
  C.wasi_cli_environment_initial_cwd(&ret)
  var lift_ret Option[string]
  if ret.is_some {
    var lift_ret_val string
    lift_ret_val = C.GoStringN(ret.val.ptr, C.int(ret.val.len))
    lift_ret.Set(lift_ret_val)
  } else {
    lift_ret.Unset()
  }
  return lift_ret
}

// Export functions from ilhan:golem-poll/api
var ilhan_golem_poll_api ExportsIlhanGolemPollApi = nil
func SetExportsIlhanGolemPollApi(i ExportsIlhanGolemPollApi) {
  ilhan_golem_poll_api = i
}
type ExportsIlhanGolemPollApi interface {
  Create(poll_data IlhanGolemPollApiPoll) Result[IlhanGolemPollApiPoll, string] 
  Respond(answer IlhanGolemPollApiResponseCreation) Result[struct{}, string] 
  Get(access_token string) Result[IlhanGolemPollApiPollData, string] 
  Close(access_token string) Result[struct{}, string] 
  Delete(access_token string) Result[struct{}, string] 
}
//export exports_ilhan_golem_poll_api_create
func ExportsIlhanGolemPollApiCreate(poll_data *C.ilhan_golem_poll_api_poll_t, ret *C.poll_result_ilhan_golem_poll_api_poll_string_t) {
  defer C.ilhan_golem_poll_api_poll_free(poll_data)
  var lift_poll_data IlhanGolemPollApiPoll
  var lift_poll_data_OwnerEmail string
  lift_poll_data_OwnerEmail = C.GoStringN(poll_data.owner_email.ptr, C.int(poll_data.owner_email.len))
  lift_poll_data.OwnerEmail = lift_poll_data_OwnerEmail
  var lift_poll_data_Title string
  lift_poll_data_Title = C.GoStringN(poll_data.title.ptr, C.int(poll_data.title.len))
  lift_poll_data.Title = lift_poll_data_Title
  var lift_poll_data_End int64
  lift_poll_data_End = int64(poll_data.end)
  lift_poll_data.End = lift_poll_data_End
  var lift_poll_data_Questions []IlhanGolemPollApiQuestion
  lift_poll_data_Questions = make([]IlhanGolemPollApiQuestion, poll_data.questions.len)
  if poll_data.questions.len > 0 {
    for lift_poll_data_Questions_i := 0; lift_poll_data_Questions_i < int(poll_data.questions.len); lift_poll_data_Questions_i++ {
      var empty_lift_poll_data_Questions C.ilhan_golem_poll_api_question_t
      lift_poll_data_Questions_ptr := *(*C.ilhan_golem_poll_api_question_t)(unsafe.Pointer(uintptr(unsafe.Pointer(poll_data.questions.ptr)) +
      uintptr(lift_poll_data_Questions_i)*unsafe.Sizeof(empty_lift_poll_data_Questions)))
      var list_lift_poll_data_Questions IlhanGolemPollApiQuestion
      var list_lift_poll_data_Questions_Text string
      list_lift_poll_data_Questions_Text = C.GoStringN(lift_poll_data_Questions_ptr.text.ptr, C.int(lift_poll_data_Questions_ptr.text.len))
      list_lift_poll_data_Questions.Text = list_lift_poll_data_Questions_Text
      var list_lift_poll_data_Questions_Answers []string
      list_lift_poll_data_Questions_Answers = make([]string, lift_poll_data_Questions_ptr.answers.len)
      if lift_poll_data_Questions_ptr.answers.len > 0 {
        for list_lift_poll_data_Questions_Answers_i := 0; list_lift_poll_data_Questions_Answers_i < int(lift_poll_data_Questions_ptr.answers.len); list_lift_poll_data_Questions_Answers_i++ {
          var empty_list_lift_poll_data_Questions_Answers C.poll_string_t
          list_lift_poll_data_Questions_Answers_ptr := *(*C.poll_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lift_poll_data_Questions_ptr.answers.ptr)) +
          uintptr(list_lift_poll_data_Questions_Answers_i)*unsafe.Sizeof(empty_list_lift_poll_data_Questions_Answers)))
          var list_list_lift_poll_data_Questions_Answers string
          list_list_lift_poll_data_Questions_Answers = C.GoStringN(list_lift_poll_data_Questions_Answers_ptr.ptr, C.int(list_lift_poll_data_Questions_Answers_ptr.len))
          list_lift_poll_data_Questions_Answers[list_lift_poll_data_Questions_Answers_i] = list_list_lift_poll_data_Questions_Answers
        }
      }
      list_lift_poll_data_Questions.Answers = list_lift_poll_data_Questions_Answers
      lift_poll_data_Questions[lift_poll_data_Questions_i] = list_lift_poll_data_Questions
    }
  }
  lift_poll_data.Questions = lift_poll_data_Questions
  result := ilhan_golem_poll_api.Create(lift_poll_data)
  var lower_result C.poll_result_ilhan_golem_poll_api_poll_string_t
  lower_result.is_err = result.IsErr()
  if result.IsOk() {
    lower_result_ptr := (*C.ilhan_golem_poll_api_poll_t)(unsafe.Pointer(&lower_result.val))
    var lower_result_val C.ilhan_golem_poll_api_poll_t
    var lower_result_val_owner_email C.poll_string_t
    
    lower_result_val_owner_email.ptr = C.CString(result.Unwrap().OwnerEmail)
    lower_result_val_owner_email.len = C.size_t(len(result.Unwrap().OwnerEmail))
    lower_result_val.owner_email = lower_result_val_owner_email
    var lower_result_val_title C.poll_string_t
    
    lower_result_val_title.ptr = C.CString(result.Unwrap().Title)
    lower_result_val_title.len = C.size_t(len(result.Unwrap().Title))
    lower_result_val.title = lower_result_val_title
    lower_result_val_end := C.int64_t(result.Unwrap().End)
    lower_result_val.end = lower_result_val_end
    var lower_result_val_questions C.poll_list_ilhan_golem_poll_api_question_t
    if len(result.Unwrap().Questions) == 0 {
      lower_result_val_questions.ptr = nil
      lower_result_val_questions.len = 0
    } else {
      var empty_lower_result_val_questions C.ilhan_golem_poll_api_question_t
      lower_result_val_questions.ptr = (*C.ilhan_golem_poll_api_question_t)(C.malloc(C.size_t(len(result.Unwrap().Questions)) * C.size_t(unsafe.Sizeof(empty_lower_result_val_questions))))
      lower_result_val_questions.len = C.size_t(len(result.Unwrap().Questions))
      for lower_result_val_questions_i := range result.Unwrap().Questions {
        lower_result_val_questions_ptr := (*C.ilhan_golem_poll_api_question_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_result_val_questions.ptr)) +
        uintptr(lower_result_val_questions_i)*unsafe.Sizeof(empty_lower_result_val_questions)))
        var lower_result_val_questions_ptr_value C.ilhan_golem_poll_api_question_t
        var lower_result_val_questions_ptr_value_text C.poll_string_t
        
        lower_result_val_questions_ptr_value_text.ptr = C.CString(result.Unwrap().Questions[lower_result_val_questions_i].Text)
        lower_result_val_questions_ptr_value_text.len = C.size_t(len(result.Unwrap().Questions[lower_result_val_questions_i].Text))
        lower_result_val_questions_ptr_value.text = lower_result_val_questions_ptr_value_text
        var lower_result_val_questions_ptr_value_answers C.poll_list_string_t
        if len(result.Unwrap().Questions[lower_result_val_questions_i].Answers) == 0 {
          lower_result_val_questions_ptr_value_answers.ptr = nil
          lower_result_val_questions_ptr_value_answers.len = 0
        } else {
          var empty_lower_result_val_questions_ptr_value_answers C.poll_string_t
          lower_result_val_questions_ptr_value_answers.ptr = (*C.poll_string_t)(C.malloc(C.size_t(len(result.Unwrap().Questions[lower_result_val_questions_i].Answers)) * C.size_t(unsafe.Sizeof(empty_lower_result_val_questions_ptr_value_answers))))
          lower_result_val_questions_ptr_value_answers.len = C.size_t(len(result.Unwrap().Questions[lower_result_val_questions_i].Answers))
          for lower_result_val_questions_ptr_value_answers_i := range result.Unwrap().Questions[lower_result_val_questions_i].Answers {
            lower_result_val_questions_ptr_value_answers_ptr := (*C.poll_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_result_val_questions_ptr_value_answers.ptr)) +
            uintptr(lower_result_val_questions_ptr_value_answers_i)*unsafe.Sizeof(empty_lower_result_val_questions_ptr_value_answers)))
            var lower_result_val_questions_ptr_value_answers_ptr_value C.poll_string_t
            
            lower_result_val_questions_ptr_value_answers_ptr_value.ptr = C.CString(result.Unwrap().Questions[lower_result_val_questions_i].Answers[lower_result_val_questions_ptr_value_answers_i])
            lower_result_val_questions_ptr_value_answers_ptr_value.len = C.size_t(len(result.Unwrap().Questions[lower_result_val_questions_i].Answers[lower_result_val_questions_ptr_value_answers_i]))
            *lower_result_val_questions_ptr_value_answers_ptr = lower_result_val_questions_ptr_value_answers_ptr_value
          }
        }
        lower_result_val_questions_ptr_value.answers = lower_result_val_questions_ptr_value_answers
        *lower_result_val_questions_ptr = lower_result_val_questions_ptr_value
      }
    }
    lower_result_val.questions = lower_result_val_questions
    *lower_result_ptr = lower_result_val
  } else {
    lower_result_ptr := (*C.poll_string_t)(unsafe.Pointer(&lower_result.val))
    var lower_result_val C.poll_string_t
    
    lower_result_val.ptr = C.CString(result.UnwrapErr())
    lower_result_val.len = C.size_t(len(result.UnwrapErr()))
    *lower_result_ptr = lower_result_val
  }
  *ret = lower_result
  
}
//export exports_ilhan_golem_poll_api_respond
func ExportsIlhanGolemPollApiRespond(answer *C.ilhan_golem_poll_api_response_creation_t, ret *C.poll_result_void_string_t) {
  defer C.ilhan_golem_poll_api_response_creation_free(answer)
  var lift_answer IlhanGolemPollApiResponseCreation
  var lift_answer_UserToken string
  lift_answer_UserToken = C.GoStringN(answer.user_token.ptr, C.int(answer.user_token.len))
  lift_answer.UserToken = lift_answer_UserToken
  var lift_answer_Answers []uint32
  lift_answer_Answers = make([]uint32, answer.answers.len)
  if answer.answers.len > 0 {
    for lift_answer_Answers_i := 0; lift_answer_Answers_i < int(answer.answers.len); lift_answer_Answers_i++ {
      var empty_lift_answer_Answers C.uint32_t
      lift_answer_Answers_ptr := *(*C.uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(answer.answers.ptr)) +
      uintptr(lift_answer_Answers_i)*unsafe.Sizeof(empty_lift_answer_Answers)))
      var list_lift_answer_Answers uint32
      list_lift_answer_Answers = uint32(lift_answer_Answers_ptr)
      lift_answer_Answers[lift_answer_Answers_i] = list_lift_answer_Answers
    }
  }
  lift_answer.Answers = lift_answer_Answers
  result := ilhan_golem_poll_api.Respond(lift_answer)
  var lower_result C.poll_result_void_string_t
  lower_result.is_err = result.IsErr()
  if result.IsOk() {
  } else {
    lower_result_ptr := (*C.poll_string_t)(unsafe.Pointer(&lower_result.val))
    var lower_result_val C.poll_string_t
    
    lower_result_val.ptr = C.CString(result.UnwrapErr())
    lower_result_val.len = C.size_t(len(result.UnwrapErr()))
    *lower_result_ptr = lower_result_val
  }
  *ret = lower_result
  
}
//export exports_ilhan_golem_poll_api_get
func ExportsIlhanGolemPollApiGet(access_token *C.poll_string_t, ret *C.poll_result_ilhan_golem_poll_api_poll_data_string_t) {
  defer C.poll_string_free(access_token)
  var lift_access_token string
  lift_access_token = C.GoStringN(access_token.ptr, C.int(access_token.len))
  result := ilhan_golem_poll_api.Get(lift_access_token)
  var lower_result C.poll_result_ilhan_golem_poll_api_poll_data_string_t
  lower_result.is_err = result.IsErr()
  if result.IsOk() {
    lower_result_ptr := (*C.ilhan_golem_poll_api_poll_data_t)(unsafe.Pointer(&lower_result.val))
    var lower_result_val C.ilhan_golem_poll_api_poll_data_t
    var lower_result_val_poll C.ilhan_golem_poll_api_poll_t
    var lower_result_val_poll_owner_email C.poll_string_t
    
    lower_result_val_poll_owner_email.ptr = C.CString(result.Unwrap().Poll.OwnerEmail)
    lower_result_val_poll_owner_email.len = C.size_t(len(result.Unwrap().Poll.OwnerEmail))
    lower_result_val_poll.owner_email = lower_result_val_poll_owner_email
    var lower_result_val_poll_title C.poll_string_t
    
    lower_result_val_poll_title.ptr = C.CString(result.Unwrap().Poll.Title)
    lower_result_val_poll_title.len = C.size_t(len(result.Unwrap().Poll.Title))
    lower_result_val_poll.title = lower_result_val_poll_title
    lower_result_val_poll_end := C.int64_t(result.Unwrap().Poll.End)
    lower_result_val_poll.end = lower_result_val_poll_end
    var lower_result_val_poll_questions C.poll_list_ilhan_golem_poll_api_question_t
    if len(result.Unwrap().Poll.Questions) == 0 {
      lower_result_val_poll_questions.ptr = nil
      lower_result_val_poll_questions.len = 0
    } else {
      var empty_lower_result_val_poll_questions C.ilhan_golem_poll_api_question_t
      lower_result_val_poll_questions.ptr = (*C.ilhan_golem_poll_api_question_t)(C.malloc(C.size_t(len(result.Unwrap().Poll.Questions)) * C.size_t(unsafe.Sizeof(empty_lower_result_val_poll_questions))))
      lower_result_val_poll_questions.len = C.size_t(len(result.Unwrap().Poll.Questions))
      for lower_result_val_poll_questions_i := range result.Unwrap().Poll.Questions {
        lower_result_val_poll_questions_ptr := (*C.ilhan_golem_poll_api_question_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_result_val_poll_questions.ptr)) +
        uintptr(lower_result_val_poll_questions_i)*unsafe.Sizeof(empty_lower_result_val_poll_questions)))
        var lower_result_val_poll_questions_ptr_value C.ilhan_golem_poll_api_question_t
        var lower_result_val_poll_questions_ptr_value_text C.poll_string_t
        
        lower_result_val_poll_questions_ptr_value_text.ptr = C.CString(result.Unwrap().Poll.Questions[lower_result_val_poll_questions_i].Text)
        lower_result_val_poll_questions_ptr_value_text.len = C.size_t(len(result.Unwrap().Poll.Questions[lower_result_val_poll_questions_i].Text))
        lower_result_val_poll_questions_ptr_value.text = lower_result_val_poll_questions_ptr_value_text
        var lower_result_val_poll_questions_ptr_value_answers C.poll_list_string_t
        if len(result.Unwrap().Poll.Questions[lower_result_val_poll_questions_i].Answers) == 0 {
          lower_result_val_poll_questions_ptr_value_answers.ptr = nil
          lower_result_val_poll_questions_ptr_value_answers.len = 0
        } else {
          var empty_lower_result_val_poll_questions_ptr_value_answers C.poll_string_t
          lower_result_val_poll_questions_ptr_value_answers.ptr = (*C.poll_string_t)(C.malloc(C.size_t(len(result.Unwrap().Poll.Questions[lower_result_val_poll_questions_i].Answers)) * C.size_t(unsafe.Sizeof(empty_lower_result_val_poll_questions_ptr_value_answers))))
          lower_result_val_poll_questions_ptr_value_answers.len = C.size_t(len(result.Unwrap().Poll.Questions[lower_result_val_poll_questions_i].Answers))
          for lower_result_val_poll_questions_ptr_value_answers_i := range result.Unwrap().Poll.Questions[lower_result_val_poll_questions_i].Answers {
            lower_result_val_poll_questions_ptr_value_answers_ptr := (*C.poll_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_result_val_poll_questions_ptr_value_answers.ptr)) +
            uintptr(lower_result_val_poll_questions_ptr_value_answers_i)*unsafe.Sizeof(empty_lower_result_val_poll_questions_ptr_value_answers)))
            var lower_result_val_poll_questions_ptr_value_answers_ptr_value C.poll_string_t
            
            lower_result_val_poll_questions_ptr_value_answers_ptr_value.ptr = C.CString(result.Unwrap().Poll.Questions[lower_result_val_poll_questions_i].Answers[lower_result_val_poll_questions_ptr_value_answers_i])
            lower_result_val_poll_questions_ptr_value_answers_ptr_value.len = C.size_t(len(result.Unwrap().Poll.Questions[lower_result_val_poll_questions_i].Answers[lower_result_val_poll_questions_ptr_value_answers_i]))
            *lower_result_val_poll_questions_ptr_value_answers_ptr = lower_result_val_poll_questions_ptr_value_answers_ptr_value
          }
        }
        lower_result_val_poll_questions_ptr_value.answers = lower_result_val_poll_questions_ptr_value_answers
        *lower_result_val_poll_questions_ptr = lower_result_val_poll_questions_ptr_value
      }
    }
    lower_result_val_poll.questions = lower_result_val_poll_questions
    lower_result_val.poll = lower_result_val_poll
    lower_result_val_open := result.Unwrap().Open
    lower_result_val.open = lower_result_val_open
    var lower_result_val_responses C.poll_list_ilhan_golem_poll_api_response_t
    if len(result.Unwrap().Responses) == 0 {
      lower_result_val_responses.ptr = nil
      lower_result_val_responses.len = 0
    } else {
      var empty_lower_result_val_responses C.ilhan_golem_poll_api_response_t
      lower_result_val_responses.ptr = (*C.ilhan_golem_poll_api_response_t)(C.malloc(C.size_t(len(result.Unwrap().Responses)) * C.size_t(unsafe.Sizeof(empty_lower_result_val_responses))))
      lower_result_val_responses.len = C.size_t(len(result.Unwrap().Responses))
      for lower_result_val_responses_i := range result.Unwrap().Responses {
        lower_result_val_responses_ptr := (*C.ilhan_golem_poll_api_response_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_result_val_responses.ptr)) +
        uintptr(lower_result_val_responses_i)*unsafe.Sizeof(empty_lower_result_val_responses)))
        var lower_result_val_responses_ptr_value C.ilhan_golem_poll_api_response_t
        var lower_result_val_responses_ptr_value_user C.poll_string_t
        
        lower_result_val_responses_ptr_value_user.ptr = C.CString(result.Unwrap().Responses[lower_result_val_responses_i].User)
        lower_result_val_responses_ptr_value_user.len = C.size_t(len(result.Unwrap().Responses[lower_result_val_responses_i].User))
        lower_result_val_responses_ptr_value.user = lower_result_val_responses_ptr_value_user
        var lower_result_val_responses_ptr_value_answers C.poll_list_u32_t
        if len(result.Unwrap().Responses[lower_result_val_responses_i].Answers) == 0 {
          lower_result_val_responses_ptr_value_answers.ptr = nil
          lower_result_val_responses_ptr_value_answers.len = 0
        } else {
          var empty_lower_result_val_responses_ptr_value_answers C.uint32_t
          lower_result_val_responses_ptr_value_answers.ptr = (*C.uint32_t)(C.malloc(C.size_t(len(result.Unwrap().Responses[lower_result_val_responses_i].Answers)) * C.size_t(unsafe.Sizeof(empty_lower_result_val_responses_ptr_value_answers))))
          lower_result_val_responses_ptr_value_answers.len = C.size_t(len(result.Unwrap().Responses[lower_result_val_responses_i].Answers))
          for lower_result_val_responses_ptr_value_answers_i := range result.Unwrap().Responses[lower_result_val_responses_i].Answers {
            lower_result_val_responses_ptr_value_answers_ptr := (*C.uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(lower_result_val_responses_ptr_value_answers.ptr)) +
            uintptr(lower_result_val_responses_ptr_value_answers_i)*unsafe.Sizeof(empty_lower_result_val_responses_ptr_value_answers)))
            lower_result_val_responses_ptr_value_answers_ptr_value := C.uint32_t(result.Unwrap().Responses[lower_result_val_responses_i].Answers[lower_result_val_responses_ptr_value_answers_i])
            *lower_result_val_responses_ptr_value_answers_ptr = lower_result_val_responses_ptr_value_answers_ptr_value
          }
        }
        lower_result_val_responses_ptr_value.answers = lower_result_val_responses_ptr_value_answers
        *lower_result_val_responses_ptr = lower_result_val_responses_ptr_value
      }
    }
    lower_result_val.responses = lower_result_val_responses
    *lower_result_ptr = lower_result_val
  } else {
    lower_result_ptr := (*C.poll_string_t)(unsafe.Pointer(&lower_result.val))
    var lower_result_val C.poll_string_t
    
    lower_result_val.ptr = C.CString(result.UnwrapErr())
    lower_result_val.len = C.size_t(len(result.UnwrapErr()))
    *lower_result_ptr = lower_result_val
  }
  *ret = lower_result
  
}
//export exports_ilhan_golem_poll_api_close
func ExportsIlhanGolemPollApiClose(access_token *C.poll_string_t, ret *C.poll_result_void_string_t) {
  defer C.poll_string_free(access_token)
  var lift_access_token string
  lift_access_token = C.GoStringN(access_token.ptr, C.int(access_token.len))
  result := ilhan_golem_poll_api.Close(lift_access_token)
  var lower_result C.poll_result_void_string_t
  lower_result.is_err = result.IsErr()
  if result.IsOk() {
  } else {
    lower_result_ptr := (*C.poll_string_t)(unsafe.Pointer(&lower_result.val))
    var lower_result_val C.poll_string_t
    
    lower_result_val.ptr = C.CString(result.UnwrapErr())
    lower_result_val.len = C.size_t(len(result.UnwrapErr()))
    *lower_result_ptr = lower_result_val
  }
  *ret = lower_result
  
}
//export exports_ilhan_golem_poll_api_delete
func ExportsIlhanGolemPollApiDelete(access_token *C.poll_string_t, ret *C.poll_result_void_string_t) {
  defer C.poll_string_free(access_token)
  var lift_access_token string
  lift_access_token = C.GoStringN(access_token.ptr, C.int(access_token.len))
  result := ilhan_golem_poll_api.Delete(lift_access_token)
  var lower_result C.poll_result_void_string_t
  lower_result.is_err = result.IsErr()
  if result.IsOk() {
  } else {
    lower_result_ptr := (*C.poll_string_t)(unsafe.Pointer(&lower_result.val))
    var lower_result_val C.poll_string_t
    
    lower_result_val.ptr = C.CString(result.UnwrapErr())
    lower_result_val.len = C.size_t(len(result.UnwrapErr()))
    *lower_result_ptr = lower_result_val
  }
  *ret = lower_result
  
}