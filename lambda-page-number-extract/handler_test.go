package main

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/textract"
)

// detectPageNumberのテスト
type TestTextract struct {
	TextractSession *session.Session
}

func (t *TestTextract) DetectDocumentText(input *textract.DetectDocumentTextInput) (*textract.DetectDocumentTextOutput, error) {
	//resp, err := t.DetectDocumentText(input)
	resp := &textract.DetectDocumentTextOutput{
		Blocks: []*textract.Block{
			{
				BlockType:       aws.String("WORD"),
				ColumnIndex:     nil,
				ColumnSpan:      nil,
				Confidence:      nil,
				EntityTypes:     nil,
				Geometry:        nil,
				Id:              nil,
				Page:            nil,
				Relationships:   nil,
				RowIndex:        nil,
				RowSpan:         nil,
				SelectionStatus: nil,
				Text:            aws.String("6"),
				TextType:        aws.String("PRINTED"),
			},
			{
				BlockType:       aws.String("WORD"),
				ColumnIndex:     nil,
				ColumnSpan:      nil,
				Confidence:      nil,
				EntityTypes:     nil,
				Geometry:        nil,
				Id:              nil,
				Page:            nil,
				Relationships:   nil,
				RowIndex:        nil,
				RowSpan:         nil,
				SelectionStatus: nil,
				Text:            aws.String("/"),
				TextType:        aws.String("PRINTED"),
			},
			{
				BlockType:       aws.String("WORD"),
				ColumnIndex:     nil,
				ColumnSpan:      nil,
				Confidence:      nil,
				EntityTypes:     nil,
				Geometry:        nil,
				Id:              nil,
				Page:            nil,
				Relationships:   nil,
				RowIndex:        nil,
				RowSpan:         nil,
				SelectionStatus: nil,
				Text:            aws.String("49"),
				TextType:        aws.String("PRINTED"),
			},
		},
		DetectDocumentTextModelVersion: aws.String("1.0"),
		DocumentMetadata: &textract.DocumentMetadata{
			Pages: aws.Int64(1),
		},
	}
	return resp, nil
}

func NewTestTextract(textractSession *session.Session) (OCRClient, error) {
	testTextractClient := TestTextract{
		TextractSession: textractSession,
	}
	return &testTextractClient, nil
}

func TestDetectPageNumber(t *testing.T) {
	expected := 6

	bucket := "web-snapshot-s3-us-east-1"
	key := "sample.png"
	region := "us-east-1"
	testTextractSession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	testTextractClient, err := NewTestTextract(testTextractSession)
	if err != nil {
		t.Errorf("error occured.\nMSG:\n\t%s", err)
	}
	simplePageNumExtracter := new(SimplePageNumExtracter)

	actual, err := detectPageNumber(testTextractClient, simplePageNumExtracter, bucket, key)
	if err != nil {
		t.Errorf("error occured.\nMSG:\n\t%s", err)
	}
	if actual != expected {
		t.Errorf("got: %v want: %v", actual, expected)
	}
}

// detectDocumentTextOutputToStringSliceのテスト
// textract.DetectDocumentTextOutputが与えられた時、それを文字列スライスにして返すこと
func TestDetectDocumentTextOutputToStringSlice(t *testing.T) {
	detectDocumentTextOutput := textract.DetectDocumentTextOutput{
		Blocks: []*textract.Block{
			{
				BlockType:       aws.String("WORD"),
				ColumnIndex:     nil,
				ColumnSpan:      nil,
				Confidence:      nil,
				EntityTypes:     nil,
				Geometry:        nil,
				Id:              nil,
				Page:            nil,
				Relationships:   nil,
				RowIndex:        nil,
				RowSpan:         nil,
				SelectionStatus: nil,
				Text:            aws.String("word"),
				TextType:        aws.String("PRINTED"),
			},
		},
		DetectDocumentTextModelVersion: aws.String("1.0"),
		DocumentMetadata: &textract.DocumentMetadata{
			Pages: aws.Int64(1),
		},
	}
	expected := []string{"word"}

	actual, err := detectDocumentTextOutputToStringSlice(&detectDocumentTextOutput)
	if err != nil {
		t.Errorf("error occured.\nMSG:\n\t%s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got: %v want: %v", actual, expected)
	}
}
