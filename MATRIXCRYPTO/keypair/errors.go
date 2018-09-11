// Copyright 2018 The MATRIX Authors 
// This file is part of the MATRIX library. 
// 
// The MATRIX library is free software: you can redistribute it and/or modify 
// it under the terms of the GNU Lesser General Public License as published by 
// the Free Software Foundation, either version 3 of the License, or 
// (at your option) any later version. 
// 
// The MATRIX library is distributed in the hope that it will be useful, 
// but WITHOUT ANY WARRANTY; without even the implied warranty of 
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
// GNU Lesser General Public License for more details. 
// 
// You should have received a copy of the GNU Lesser General Public License 
// along with the MATRIX library. If not, see <http://www.gnu.org/licenses/>. 

package keypair

type EncryptError struct {
	detail string
}

func (e *EncryptError) Error() string {
	return "encrypt private key error: " + e.detail
}

func NewEncryptError(msg string) *EncryptError {
	return &EncryptError{detail: msg}
}

type DecryptError EncryptError

func (e *DecryptError) Error() string {
	return "decrypt private key error: " + e.detail
}

func NewDecryptError(msg string) *DecryptError {
	return &DecryptError{detail: msg}
}
