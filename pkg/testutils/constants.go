// Copyright 2025 César Nieto Sánchez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package testutils

// TestToleranceRelative defines the relative tolerance used for comparing floating-point numbers in tests.
// It represents the maximum allowed relative difference between the expected and actual values for the test to pass.
// A value of 1e-12 ensures that the comparison is precise up to 12 decimal places.
const TestToleranceRelative = 1e-15

// TestToleranceAbsolute defines the absolute tolerance used for comparing floating-point numbers in tests.
// It represents the maximum allowed absolute difference between the expected and actual values for the test to pass.
// A value of 1e-9 ensures that the comparison allows small discrepancies in the least significant digits, useful for large numbers.
const TestToleranceAbsolute = 1e-9
