package animelayer_comparator

import (
	animelayer_model "collector/pkg/animelayer/model"
	"slices"
)

func isCompletedToString(completed bool) string {
	if completed {
		return "Completed"
	}

	return "Incomplete"
}

func compareField[T any](newItem, oldItem T, fieldGetter func(e T) string, name string) *animelayer_model.Difference {

	oldValue := fieldGetter(oldItem)
	newValue := fieldGetter(newItem)

	if oldValue == newValue {
		return nil
	}

	return &animelayer_model.Difference{
		Name:     name,
		OldValue: oldValue,
		NewValue: newValue,
	}
}

type fieldComparer[T any] struct {
	Name     string
	GetValue func(e *T) string
}

var (
	items_fields = []fieldComparer[animelayer_model.Item]{
		{
			GetValue: func(e *animelayer_model.Item) string { return e.GUID },
			Name:     "Title",
		},
		{
			GetValue: func(e *animelayer_model.Item) string { return isCompletedToString(e.Completed) },
			Name:     "IsCompleted",
		},
	}
	description_fields = []fieldComparer[animelayer_model.ItemDescription]{
		{
			GetValue: func(e *animelayer_model.ItemDescription) string { return e.GUID },
			Name:     "Title",
		},
		{
			GetValue: func(e *animelayer_model.ItemDescription) string { return e.TorrentFilesSize },
			Name:     "TorrentFilesSize",
		},
		{
			GetValue: func(e *animelayer_model.ItemDescription) string { return e.RefImagePreview },
			Name:     "RefImagePreview",
		},
		{
			GetValue: func(e *animelayer_model.ItemDescription) string { return e.RefImageCover },
			Name:     "RefImageCover",
		},
		{
			GetValue: func(e *animelayer_model.ItemDescription) string { return e.UpdatedDate },
			Name:     "UpdatedDate",
		},
		{
			GetValue: func(e *animelayer_model.ItemDescription) string { return e.CreatedDate },
			Name:     "CreatedDate",
		},
		{
			GetValue: func(e *animelayer_model.ItemDescription) string { return e.LastCheckedDate },
			Name:     "LastCheckedDate",
		},
	}
)

func CompareItems(newItem, oldItem *animelayer_model.Item) []animelayer_model.Difference {

	diff := make([]animelayer_model.Difference, 0, 2)

	for _, field := range items_fields {

		d := compareField(newItem, oldItem, field.GetValue, field.Name)
		if d != nil {
			diff = append(diff, *d)
		}

	}

	return diff
}

func CompareDescriptions(newItem, oldItem *animelayer_model.ItemDescription) []animelayer_model.Difference {

	diff := make([]animelayer_model.Difference, 0, 2)

	for _, field := range description_fields {

		d := compareField(newItem, oldItem, field.GetValue, field.Name)
		if d != nil {
			diff = append(diff, *d)
		}

	}

	oldNotesIndexes := make([]int, len(oldItem.Descriptions))
	for i := range oldNotesIndexes {
		oldNotesIndexes[i] = i
	}

	for _, field := range newItem.Descriptions {

		key := field.Key
		newValue := field.Value

		oldIndex := slices.IndexFunc(
			oldItem.Descriptions,
			func(p animelayer_model.DescriptionPoint) bool {
				return p.Key == key
			},
		)

		if oldIndex == -1 {
			diff = append(diff,
				animelayer_model.Difference{
					Name:     key,
					OldValue: "",
					NewValue: newValue,
				},
			)
			continue
		}

		oldNotesIndexes = slices.DeleteFunc(oldNotesIndexes, func(e int) bool { return e == oldIndex })

		oldValue := oldItem.Descriptions[oldIndex].Value
		if oldValue != newValue {
			diff = append(diff,
				animelayer_model.Difference{
					Name:     key,
					OldValue: oldValue,
					NewValue: newValue,
				},
			)
		}
	}

	for _, i := range oldNotesIndexes {
		diff = append(diff,
			animelayer_model.Difference{
				Name:     oldItem.Descriptions[i].Key,
				OldValue: oldItem.Descriptions[i].Value,
				NewValue: "",
			},
		)

	}

	return diff
}
