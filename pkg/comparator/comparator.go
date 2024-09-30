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
		/*
			{
				GetValue: func(e *animelayer_model.Item) string { return e.Identifier },
				Name:     "Title",
			},
		*/
		{
			GetValue: func(e *animelayer_model.Item) string { return isCompletedToString(e.IsCompleted) },
			Name:     "IsCompleted",
		},
	}
	description_fields = []fieldComparer[animelayer_model.Description]{
		/*
				{
				GetValue: func(e *animelayer_model.Description) string { return e.Identifier },
				Name:     "Title",
			},
		*/
		{
			GetValue: func(e *animelayer_model.Description) string { return e.TorrentFilesSize },
			Name:     "TorrentFilesSize",
		},
		{
			GetValue: func(e *animelayer_model.Description) string { return e.RefImagePreview },
			Name:     "RefImagePreview",
		},
		{
			GetValue: func(e *animelayer_model.Description) string { return e.RefImageCover },
			Name:     "RefImageCover",
		},
		{
			GetValue: func(e *animelayer_model.Description) string { return e.UpdatedDate },
			Name:     "UpdatedDate",
		},
		{
			GetValue: func(e *animelayer_model.Description) string { return e.CreatedDate },
			Name:     "CreatedDate",
		},
		{
			GetValue: func(e *animelayer_model.Description) string { return e.LastCheckedDate },
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

type NotesChangedIdexes struct {
	IndexNewAdded   []int
	IndexOldRemoved []int
	IndexNewChanged []int
}

func CompareDescriptions(newItem, oldItem *animelayer_model.Description) ([]animelayer_model.Difference, *NotesChangedIdexes) {

	notesUpdate := &NotesChangedIdexes{}
	diff := make([]animelayer_model.Difference, 0, 2)

	for _, field := range description_fields {

		d := compareField(newItem, oldItem, field.GetValue, field.Name)
		if d != nil {
			diff = append(diff, *d)
		}

	}

	oldNotesIndexes := make([]int, len(oldItem.Notes))
	for i := range oldNotesIndexes {
		oldNotesIndexes[i] = i
	}

	for i, field := range newItem.Notes {

		key := field.Name
		newValue := field.Text

		oldIndex := slices.IndexFunc(
			oldItem.Notes,
			func(p animelayer_model.DescriptionNote) bool {
				return p.Name == key
			},
		)

		if oldIndex == -1 {
			notesUpdate.IndexNewAdded = append(notesUpdate.IndexNewAdded, i)
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

		oldValue := oldItem.Notes[oldIndex].Text
		if oldValue != newValue {
			notesUpdate.IndexNewChanged = append(notesUpdate.IndexNewChanged, i)
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
		notesUpdate.IndexOldRemoved = append(notesUpdate.IndexOldRemoved, i)
		diff = append(diff,
			animelayer_model.Difference{
				Name:     oldItem.Notes[i].Name,
				OldValue: oldItem.Notes[i].Text,
				NewValue: "",
			},
		)

	}

	return diff, notesUpdate
}
