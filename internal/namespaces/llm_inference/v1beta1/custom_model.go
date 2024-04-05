package llm_inference

import (
	"github.com/scaleway/scaleway-cli/v2/internal/human"
	llm_inference "github.com/scaleway/scaleway-sdk-go/api/llm_inference/v1beta1"
)

func ListModelMarshalerFunc(i interface{}, opt *human.MarshalOpt) (string, error) {
	type tmp []*llm_inference.Model
	model := tmp(i.([]*llm_inference.Model))
	opt.Fields = []*human.MarshalFieldOpt{
		{
			FieldName: "ID",
			Label:     "ID",
		},
		{
			FieldName: "Name",
			Label:     "Name",
		},
		{
			FieldName: "Provider",
			Label:     "Provider",
		},
		{
			FieldName: "Tags",
			Label:     "Tags",
		},
	}
	str, err := human.Marshal(model, opt)
	if err != nil {
		return "", err
	}
	return str, nil
}
