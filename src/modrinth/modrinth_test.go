/*
 * modrinth_test.go - To test the high level Modrinth API
 *
 * Copyright (C) 2026 Mrborghini
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <https://www.gnu.org/licenses/>.
 */

package modrinth

import (
	"testing"
)

func TestInfoRequest(t *testing.T) {
	m := New()
	result, err := m.info.Get()
	if err != nil {
		t.Fatalf("info.get() returned an error: %v", err)
	}
	if result == nil {
		t.Fatal("info.get() returned nil result")
	}
	if result.Version == "" {
		t.Fatal("info.get().Version is empty")
	}
}
