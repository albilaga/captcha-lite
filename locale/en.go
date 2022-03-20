package locale

var EN map[Message]string = map[Message]string{
	MessageWelcome:                "Hi, {{user}}!\n\nWelcome to {{group}}. Make sure you read pinned message first.",
	MessageKick:                   "{{user}} has been kicked because he didn't complete the captcha.",
	MessageJoin:                   "Hi, {{user}}!\n\nBefore you continue, please complete this captcha. You have 1 minute from now.\n\n{{captcha}}",
	MessageWrongAnswerLettersOnly: "Wrong answer. Only letters are allowed. You have {{remaining}} seconds left to complete.",
	MessageWrongAnswer:            "Wrong answer, please try again. You have {{remaining}} seconds left to complete.",
	MessageNonText:                "Hi, {{user}}. Complete the captcha first. You have {{remaining}} seconds left.",
}
