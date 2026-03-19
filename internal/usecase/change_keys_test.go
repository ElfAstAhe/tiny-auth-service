package usecase

import (
	"context"
	"errors"
	"testing"

	libdbmocks "github.com/ElfAstAhe/go-service-template/pkg/db/mocks"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	libhlpmocks "github.com/ElfAstAhe/go-service-template/pkg/helper/mocks"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestChangeKeys_TableDriven(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(mRepo *mocks.MockUserRepository, mKeys *libhlpmocks.MockRSAKeys)
		wantErr bool
		errType interface{} // Для проверки типа ошибки (через errors.As)
	}{
		{
			name: "Success",
			args: args{userID: "user-1"},
			prepare: func(mRepo *mocks.MockUserRepository, mKeys *libhlpmocks.MockRSAKeys) {
				mRepo.On("Find", mock.Anything, "user-1").Return(&domain.User{ID: "user-1"}, nil)
				mKeys.On("Generate").Return("priv-key", "pub-key", nil)
				mRepo.On("Change", mock.Anything, mock.Anything).Return(&domain.User{}, nil)
			},
			wantErr: false,
		},
		{
			name: "Validation error - empty ID",
			args: args{userID: ""},
			prepare: func(mRepo *mocks.MockUserRepository, mKeys *libhlpmocks.MockRSAKeys) {
				// Ничего не вызывается
			},
			wantErr: true,
		},
		{
			name: "User not found",
			args: args{userID: "404"},
			prepare: func(mRepo *mocks.MockUserRepository, mKeys *libhlpmocks.MockRSAKeys) {
				mRepo.On("Find", mock.Anything, "404").Return(nil, &errs.DalNotFoundError{})
			},
			wantErr: true,
		},
		{
			name: "RSA generation failure",
			args: args{userID: "user-1"},
			prepare: func(mRepo *mocks.MockUserRepository, mKeys *libhlpmocks.MockRSAKeys) {
				mRepo.On("Find", mock.Anything, "user-1").Return(&domain.User{ID: "user-1"}, nil)
				mKeys.On("Generate").Return("", "", errors.New("entropy error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(mocks.MockUserRepository)
			mockTM := new(libdbmocks.MockTransactionManager)
			mockKeys := new(libhlpmocks.MockRSAKeys) // Рекомендую сделать интерфейс для RSAKeysHelper

			uc := NewChangeKeysUseCase(mockKeys, mockTM, mockRepo)
			tt.prepare(mockRepo, mockKeys)

			// Act
			priv, pub, err := uc.ChangeKeys(context.Background(), tt.args.userID)

			// Assert
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, priv)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, priv)
				assert.NotEmpty(t, pub)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
