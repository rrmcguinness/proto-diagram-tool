/*
 * Copyright 2022 Google, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proto

import "fmt"

// Logger is a simplified logger for making the code and output readable.
type Logger struct {
	debug bool
}

func (l Logger) Debug(in string) {
	if l.debug {
		fmt.Printf("DEBUG: %s\n", in)
	}
}

func (l Logger) Debugf(in string, args ...any) {
	l.Debug(fmt.Sprintf(in, args...))
}

func (l Logger) Error(in string) {
	fmt.Printf("ERROR: %s\n", in)
}

func (l Logger) Errorf(in string, args ...any) {
	l.Error(fmt.Sprintf(in, args...))
}

func (l Logger) Info(in string) {
	fmt.Printf("INFO: %s\n", in)
}

func (l Logger) Infof(in string, args ...any) {
	l.Info(fmt.Sprintf(in, args...))
}