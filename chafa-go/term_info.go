// Copyright (C) 2018-2025 Hans Petter Jansson
//
// This file is part of Chafa, a program that shows pictures on text terminals.
//
// Chafa is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Chafa is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Chafa.  If not, see <http://www.gnu.org/licenses/>.

package chafa

// TermInfo represents terminal information and capabilities.
type TermInfo struct {
	// TODO: Internal structure to be implemented
}

// NewTermInfo creates a new TermInfo.
func NewTermInfo() *TermInfo {
	return &TermInfo{}
}

// Copy creates a copy of the TermInfo.
func (ti *TermInfo) Copy() *TermInfo {
	if ti == nil {
		return nil
	}
	copy := *ti
	return &copy
}
