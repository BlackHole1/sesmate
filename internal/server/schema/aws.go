package schema

// https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_SendEmail.html#API_SendEmail_RequestSyntax

// AWSRawMessage https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_RawMessage.html
type AWSRawMessage struct {
	Data string `binding:"required"`
}

// AWSContent https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_Content.html
type AWSContent struct {
	Charset string `binding:"omitempty"`
	Data    string `binding:"required"`
}

// AWSBody https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_Body.html
type AWSBody struct {
	Html *AWSContent `binding:"omitempty"`
	Text *AWSContent `binding:"omitempty"`
}

// AWSMessage https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_Message.html
type AWSMessage struct {
	Body    *AWSBody    `binding:"required"`
	Subject *AWSContent `binding:"required"`
}

// AWSTemplate https://docs.aws.amazon.com/zh_cn/ses/latest/APIReference-V2/API_Template.html
type AWSTemplate struct {
	TemplateArn  string `binding:"omitempty"`
	TemplateData string `binding:"omitempty,max=262144"`
	TemplateName string `binding:"omitempty,min=1"`
}

// AWSEmailContent https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_EmailContent.html
type AWSEmailContent struct {
	Raw      *AWSRawMessage `binding:"omitempty"`
	Simple   *AWSMessage    `binding:"omitempty"`
	Template *AWSTemplate   `binding:"omitempty"`
}

// AWSDestination https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_Destination.html
type AWSDestination struct {
	BccAddresses []string `binding:"omitempty"`
	CcAddresses  []string `binding:"omitempty"`
	ToAddresses  []string `binding:"omitempty"`
}

// AWSMessageTag  https://docs.aws.amazon.com/zh_cn/ses/latest/APIReference-V2/API_MessageTag.html
type AWSMessageTag struct {
	Name  string `binding:"required,allowedChars,max=256"`
	Value string `binding:"required,allowedChars,max=256"`
}

// AWSListManagementOptions https://docs.aws.amazon.com/zh_cn/ses/latest/APIReference-V2/API_ListManagementOptions.html
type AWSListManagementOptions struct {
	ContactListName string `binding:"required"`
	TopicName       string `binding:"omitempty"`
}
