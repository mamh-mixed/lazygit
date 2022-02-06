package context

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/context/traits"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type StashContext struct {
	*StashViewModel
	*ListContextTrait
}

var _ types.IListContext = (*StashContext)(nil)

func NewStashContext(
	getModel func() []*models.StashEntry,
	view *gocui.View,
	getDisplayStrings func(startIdx int, length int) [][]string,

	onFocus func(...types.OnFocusOpts) error,
	onRenderToMain func(...types.OnFocusOpts) error,
	onFocusLost func() error,

	c *types.ControllerCommon,
) *StashContext {
	viewModel := NewStashViewModel(getModel)

	return &StashContext{
		StashViewModel: viewModel,
		ListContextTrait: &ListContextTrait{
			Context: NewSimpleContext(NewBaseContext(NewBaseContextOpts{
				ViewName:   "stash",
				WindowName: "stash",
				Key:        STASH_CONTEXT_KEY,
				Kind:       types.SIDE_CONTEXT,
				Focusable:  true,
			}), ContextCallbackOpts{
				OnFocus:        onFocus,
				OnFocusLost:    onFocusLost,
				OnRenderToMain: onRenderToMain,
			}),
			list:              viewModel,
			viewTrait:         NewViewTrait(view),
			getDisplayStrings: getDisplayStrings,
			c:                 c,
		},
	}
}

func (self *StashContext) GetSelectedItemId() string {
	item := self.GetSelectedStashEntry()
	if item == nil {
		return ""
	}

	return item.ID()
}

type StashViewModel struct {
	*traits.ListCursor
	getModel func() []*models.StashEntry
}

func NewStashViewModel(getModel func() []*models.StashEntry) *StashViewModel {
	self := &StashViewModel{
		getModel: getModel,
	}

	self.ListCursor = traits.NewListCursor(self)

	return self
}

func (self *StashViewModel) GetItemsLength() int {
	return len(self.getModel())
}

func (self *StashViewModel) GetSelectedStashEntry() *models.StashEntry {
	if self.GetItemsLength() == 0 {
		return nil
	}

	return self.getModel()[self.GetSelectedLineIdx()]
}
