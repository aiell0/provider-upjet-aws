/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AbortStatementInitParameters struct {

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message.
	Message []MessageInitParameters `json:"message,omitempty" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card.
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type AbortStatementObservation struct {

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message.
	Message []MessageObservation `json:"message,omitempty" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card.
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type AbortStatementParameters struct {

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message.
	// +kubebuilder:validation:Optional
	Message []MessageParameters `json:"message" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card.
	// +kubebuilder:validation:Optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type BotInitParameters struct {

	// The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
	AbortStatement []AbortStatementInitParameters `json:"abortStatement,omitempty" tf:"abort_statement,omitempty"`

	// By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the Amazon Lex FAQ and the Amazon Lex PutBot API Docs.
	ChildDirected *bool `json:"childDirected,omitempty" tf:"child_directed,omitempty"`

	// The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
	ClarificationPrompt []ClarificationPromptInitParameters `json:"clarificationPrompt,omitempty" tf:"clarification_prompt,omitempty"`

	// Determines if a new bot version is created when the initial resource is created and on each update. Defaults to false.
	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version,omitempty"`

	// A description of the bot. Must be less than or equal to 200 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is false.
	DetectSentiment *bool `json:"detectSentiment,omitempty" tf:"detect_sentiment,omitempty"`

	// Set to true to enable access to natural language understanding improvements. When you set the enable_model_improvements parameter to true you can use the nlu_intent_confidence_threshold parameter to configure confidence scores. For more information, see Confidence Scores. You can only set the enable_model_improvements parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the Amazon Lex Bot PutBot API Docs.
	EnableModelImprovements *bool `json:"enableModelImprovements,omitempty" tf:"enable_model_improvements,omitempty"`

	// The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is 300. Must be a number between 60 and 86400 (inclusive).
	IdleSessionTTLInSeconds *float64 `json:"idleSessionTtlInSeconds,omitempty" tf:"idle_session_ttl_in_seconds,omitempty"`

	// A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 250 Intent objects.
	Intent []IntentInitParameters `json:"intent,omitempty" tf:"intent,omitempty"`

	// Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see Amazon Lex Bot PutBot API Docs. Default is en-US.
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see Amazon Lex Bot PutBot API Docs This value requires enable_model_improvements to be set to true and the default is 0. Must be a float between 0 and 1.
	NluIntentConfidenceThreshold *float64 `json:"nluIntentConfidenceThreshold,omitempty" tf:"nlu_intent_confidence_threshold,omitempty"`

	// If you set the process_behavior element to BUILD, Amazon Lex builds the bot so that it can be run. If you set the element to SAVE Amazon Lex saves the bot, but doesn't build it. Default is SAVE.
	ProcessBehavior *string `json:"processBehavior,omitempty" tf:"process_behavior,omitempty"`

	// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see Available Voices in the Amazon Polly Developer Guide.
	VoiceID *string `json:"voiceId,omitempty" tf:"voice_id,omitempty"`
}

type BotObservation struct {

	// The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
	AbortStatement []AbortStatementObservation `json:"abortStatement,omitempty" tf:"abort_statement,omitempty"`

	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Checksum identifying the version of the bot that was created. The checksum is not
	// included as an argument because the resource will add it automatically when updating the bot.
	Checksum *string `json:"checksum,omitempty" tf:"checksum,omitempty"`

	// By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the Amazon Lex FAQ and the Amazon Lex PutBot API Docs.
	ChildDirected *bool `json:"childDirected,omitempty" tf:"child_directed,omitempty"`

	// The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
	ClarificationPrompt []ClarificationPromptObservation `json:"clarificationPrompt,omitempty" tf:"clarification_prompt,omitempty"`

	// Determines if a new bot version is created when the initial resource is created and on each update. Defaults to false.
	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version,omitempty"`

	// The date when the bot version was created.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// A description of the bot. Must be less than or equal to 200 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is false.
	DetectSentiment *bool `json:"detectSentiment,omitempty" tf:"detect_sentiment,omitempty"`

	// Set to true to enable access to natural language understanding improvements. When you set the enable_model_improvements parameter to true you can use the nlu_intent_confidence_threshold parameter to configure confidence scores. For more information, see Confidence Scores. You can only set the enable_model_improvements parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the Amazon Lex Bot PutBot API Docs.
	EnableModelImprovements *bool `json:"enableModelImprovements,omitempty" tf:"enable_model_improvements,omitempty"`

	// If status is FAILED, Amazon Lex provides the reason that it failed to build the bot.
	FailureReason *string `json:"failureReason,omitempty" tf:"failure_reason,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is 300. Must be a number between 60 and 86400 (inclusive).
	IdleSessionTTLInSeconds *float64 `json:"idleSessionTtlInSeconds,omitempty" tf:"idle_session_ttl_in_seconds,omitempty"`

	// A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 250 Intent objects.
	Intent []IntentObservation `json:"intent,omitempty" tf:"intent,omitempty"`

	// The date when the $LATEST version of this bot was updated.
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`

	// Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see Amazon Lex Bot PutBot API Docs. Default is en-US.
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see Amazon Lex Bot PutBot API Docs This value requires enable_model_improvements to be set to true and the default is 0. Must be a float between 0 and 1.
	NluIntentConfidenceThreshold *float64 `json:"nluIntentConfidenceThreshold,omitempty" tf:"nlu_intent_confidence_threshold,omitempty"`

	// If you set the process_behavior element to BUILD, Amazon Lex builds the bot so that it can be run. If you set the element to SAVE Amazon Lex saves the bot, but doesn't build it. Default is SAVE.
	ProcessBehavior *string `json:"processBehavior,omitempty" tf:"process_behavior,omitempty"`

	// When you send a request to create or update a bot, Amazon Lex sets the status response
	// element to BUILDING. After Amazon Lex builds the bot, it sets status to READY. If Amazon Lex can't
	// build the bot, it sets status to FAILED. Amazon Lex returns the reason for the failure in the
	// failure_reason response element.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The version of the bot.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see Available Voices in the Amazon Polly Developer Guide.
	VoiceID *string `json:"voiceId,omitempty" tf:"voice_id,omitempty"`
}

type BotParameters struct {

	// The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
	// +kubebuilder:validation:Optional
	AbortStatement []AbortStatementParameters `json:"abortStatement,omitempty" tf:"abort_statement,omitempty"`

	// By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the Amazon Lex FAQ and the Amazon Lex PutBot API Docs.
	// +kubebuilder:validation:Optional
	ChildDirected *bool `json:"childDirected,omitempty" tf:"child_directed,omitempty"`

	// The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
	// +kubebuilder:validation:Optional
	ClarificationPrompt []ClarificationPromptParameters `json:"clarificationPrompt,omitempty" tf:"clarification_prompt,omitempty"`

	// Determines if a new bot version is created when the initial resource is created and on each update. Defaults to false.
	// +kubebuilder:validation:Optional
	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version,omitempty"`

	// A description of the bot. Must be less than or equal to 200 characters in length.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is false.
	// +kubebuilder:validation:Optional
	DetectSentiment *bool `json:"detectSentiment,omitempty" tf:"detect_sentiment,omitempty"`

	// Set to true to enable access to natural language understanding improvements. When you set the enable_model_improvements parameter to true you can use the nlu_intent_confidence_threshold parameter to configure confidence scores. For more information, see Confidence Scores. You can only set the enable_model_improvements parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the Amazon Lex Bot PutBot API Docs.
	// +kubebuilder:validation:Optional
	EnableModelImprovements *bool `json:"enableModelImprovements,omitempty" tf:"enable_model_improvements,omitempty"`

	// The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is 300. Must be a number between 60 and 86400 (inclusive).
	// +kubebuilder:validation:Optional
	IdleSessionTTLInSeconds *float64 `json:"idleSessionTtlInSeconds,omitempty" tf:"idle_session_ttl_in_seconds,omitempty"`

	// A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 250 Intent objects.
	// +kubebuilder:validation:Optional
	Intent []IntentParameters `json:"intent,omitempty" tf:"intent,omitempty"`

	// Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see Amazon Lex Bot PutBot API Docs. Default is en-US.
	// +kubebuilder:validation:Optional
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see Amazon Lex Bot PutBot API Docs This value requires enable_model_improvements to be set to true and the default is 0. Must be a float between 0 and 1.
	// +kubebuilder:validation:Optional
	NluIntentConfidenceThreshold *float64 `json:"nluIntentConfidenceThreshold,omitempty" tf:"nlu_intent_confidence_threshold,omitempty"`

	// If you set the process_behavior element to BUILD, Amazon Lex builds the bot so that it can be run. If you set the element to SAVE Amazon Lex saves the bot, but doesn't build it. Default is SAVE.
	// +kubebuilder:validation:Optional
	ProcessBehavior *string `json:"processBehavior,omitempty" tf:"process_behavior,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see Available Voices in the Amazon Polly Developer Guide.
	// +kubebuilder:validation:Optional
	VoiceID *string `json:"voiceId,omitempty" tf:"voice_id,omitempty"`
}

type ClarificationPromptInitParameters struct {

	// The number of times to prompt the user for information.
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message.
	Message []ClarificationPromptMessageInitParameters `json:"message,omitempty" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card.
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type ClarificationPromptMessageInitParameters struct {

	// The text of the message.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The content type of the message string.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response.
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type ClarificationPromptMessageObservation struct {

	// The text of the message.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The content type of the message string.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response.
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type ClarificationPromptMessageParameters struct {

	// The text of the message.
	// +kubebuilder:validation:Optional
	Content *string `json:"content" tf:"content,omitempty"`

	// The content type of the message string.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response.
	// +kubebuilder:validation:Optional
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type ClarificationPromptObservation struct {

	// The number of times to prompt the user for information.
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message.
	Message []ClarificationPromptMessageObservation `json:"message,omitempty" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card.
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type ClarificationPromptParameters struct {

	// The number of times to prompt the user for information.
	// +kubebuilder:validation:Optional
	MaxAttempts *float64 `json:"maxAttempts" tf:"max_attempts,omitempty"`

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message.
	// +kubebuilder:validation:Optional
	Message []ClarificationPromptMessageParameters `json:"message" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card.
	// +kubebuilder:validation:Optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type IntentInitParameters struct {

	// The name of the intent. Must be less than or equal to 100 characters in length.
	IntentName *string `json:"intentName,omitempty" tf:"intent_name,omitempty"`

	// The version of the intent. Must be less than or equal to 64 characters in length.
	IntentVersion *string `json:"intentVersion,omitempty" tf:"intent_version,omitempty"`
}

type IntentObservation struct {

	// The name of the intent. Must be less than or equal to 100 characters in length.
	IntentName *string `json:"intentName,omitempty" tf:"intent_name,omitempty"`

	// The version of the intent. Must be less than or equal to 64 characters in length.
	IntentVersion *string `json:"intentVersion,omitempty" tf:"intent_version,omitempty"`
}

type IntentParameters struct {

	// The name of the intent. Must be less than or equal to 100 characters in length.
	// +kubebuilder:validation:Optional
	IntentName *string `json:"intentName" tf:"intent_name,omitempty"`

	// The version of the intent. Must be less than or equal to 64 characters in length.
	// +kubebuilder:validation:Optional
	IntentVersion *string `json:"intentVersion" tf:"intent_version,omitempty"`
}

type MessageInitParameters struct {

	// The text of the message.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The content type of the message string.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response.
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type MessageObservation struct {

	// The text of the message.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The content type of the message string.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response.
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type MessageParameters struct {

	// The text of the message.
	// +kubebuilder:validation:Optional
	Content *string `json:"content" tf:"content,omitempty"`

	// The content type of the message string.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response.
	// +kubebuilder:validation:Optional
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

// BotSpec defines the desired state of Bot
type BotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider BotInitParameters `json:"initProvider,omitempty"`
}

// BotStatus defines the observed state of Bot.
type BotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bot is the Schema for the Bots API. Provides an Amazon Lex bot resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Bot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.abortStatement) || has(self.initProvider.abortStatement)",message="abortStatement is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.childDirected) || has(self.initProvider.childDirected)",message="childDirected is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.intent) || has(self.initProvider.intent)",message="intent is a required parameter"
	Spec   BotSpec   `json:"spec"`
	Status BotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotList contains a list of Bots
type BotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bot `json:"items"`
}

// Repository type metadata.
var (
	Bot_Kind             = "Bot"
	Bot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bot_Kind}.String()
	Bot_KindAPIVersion   = Bot_Kind + "." + CRDGroupVersion.String()
	Bot_GroupVersionKind = CRDGroupVersion.WithKind(Bot_Kind)
)

func init() {
	SchemeBuilder.Register(&Bot{}, &BotList{})
}
