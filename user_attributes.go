/* 
 * ProcessMaker API
 *
 * This ProcessMaker I/O API provides access to a BPMN 2.0 compliant workflow engine API that is designed to be used as a microservice to support enterprise cloud applications. The current Alpha 1.0 version supports most of the descriptive classes of the BPMN 2.0 specification.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: support@processmaker.io
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pmio

type UserAttributes struct {

	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`

	Firstname string `json:"firstname,omitempty"`

	Lastname string `json:"lastname,omitempty"`

	Email string `json:"email,omitempty"`

	ExpiresAt string `json:"expires_at,omitempty"`

	Status string `json:"status,omitempty"`

	Country string `json:"country,omitempty"`

	City string `json:"city,omitempty"`

	Location string `json:"location,omitempty"`

	Address string `json:"address,omitempty"`

	Phone string `json:"phone,omitempty"`

	Fax string `json:"fax,omitempty"`

	Cellular string `json:"cellular,omitempty"`

	ZipCode string `json:"zip_code,omitempty"`

	Position string `json:"position,omitempty"`

	Resume string `json:"resume,omitempty"`

	BirthdayAt string `json:"birthday_at,omitempty"`

	BookmarkStartCases string `json:"bookmark_start_cases,omitempty"`

	TimeZone string `json:"time_zone,omitempty"`

	DefaultLang string `json:"default_lang,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`

	Clients []int32 `json:"clients,omitempty"`
}
