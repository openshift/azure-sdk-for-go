// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package computervision

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/computervision"

type BaseClient = original.BaseClient
type AzureRegions = original.AzureRegions

const (
	Australiaeast  AzureRegions = original.Australiaeast
	Brazilsouth    AzureRegions = original.Brazilsouth
	Eastasia       AzureRegions = original.Eastasia
	Eastus         AzureRegions = original.Eastus
	Eastus2        AzureRegions = original.Eastus2
	Northeurope    AzureRegions = original.Northeurope
	Southcentralus AzureRegions = original.Southcentralus
	Southeastasia  AzureRegions = original.Southeastasia
	Westcentralus  AzureRegions = original.Westcentralus
	Westeurope     AzureRegions = original.Westeurope
	Westus         AzureRegions = original.Westus
	Westus2        AzureRegions = original.Westus2
)

type Details = original.Details

const (
	Celebrities Details = original.Celebrities
	Landmarks   Details = original.Landmarks
)

type DomainModels = original.DomainModels

const (
	DomainModelsCelebrities DomainModels = original.DomainModelsCelebrities
	DomainModelsLandmarks   DomainModels = original.DomainModelsLandmarks
)

type ErrorCodes = original.ErrorCodes

const (
	BadArgument               ErrorCodes = original.BadArgument
	FailedToProcess           ErrorCodes = original.FailedToProcess
	InternalServerError       ErrorCodes = original.InternalServerError
	InvalidDetails            ErrorCodes = original.InvalidDetails
	InvalidImageFormat        ErrorCodes = original.InvalidImageFormat
	InvalidImageSize          ErrorCodes = original.InvalidImageSize
	InvalidImageURL           ErrorCodes = original.InvalidImageURL
	NotSupportedImage         ErrorCodes = original.NotSupportedImage
	NotSupportedLanguage      ErrorCodes = original.NotSupportedLanguage
	NotSupportedVisualFeature ErrorCodes = original.NotSupportedVisualFeature
	StorageException          ErrorCodes = original.StorageException
	Timeout                   ErrorCodes = original.Timeout
	Unspecified               ErrorCodes = original.Unspecified
)

type Gender = original.Gender

const (
	Female Gender = original.Female
	Male   Gender = original.Male
)

type Language1 = original.Language1

const (
	En Language1 = original.En
	Zh Language1 = original.Zh
)

type OcrLanguages = original.OcrLanguages

const (
	OcrLanguagesAr     OcrLanguages = original.OcrLanguagesAr
	OcrLanguagesCs     OcrLanguages = original.OcrLanguagesCs
	OcrLanguagesDa     OcrLanguages = original.OcrLanguagesDa
	OcrLanguagesDe     OcrLanguages = original.OcrLanguagesDe
	OcrLanguagesEl     OcrLanguages = original.OcrLanguagesEl
	OcrLanguagesEn     OcrLanguages = original.OcrLanguagesEn
	OcrLanguagesEs     OcrLanguages = original.OcrLanguagesEs
	OcrLanguagesFi     OcrLanguages = original.OcrLanguagesFi
	OcrLanguagesFr     OcrLanguages = original.OcrLanguagesFr
	OcrLanguagesHu     OcrLanguages = original.OcrLanguagesHu
	OcrLanguagesIt     OcrLanguages = original.OcrLanguagesIt
	OcrLanguagesJa     OcrLanguages = original.OcrLanguagesJa
	OcrLanguagesKo     OcrLanguages = original.OcrLanguagesKo
	OcrLanguagesNb     OcrLanguages = original.OcrLanguagesNb
	OcrLanguagesNl     OcrLanguages = original.OcrLanguagesNl
	OcrLanguagesPl     OcrLanguages = original.OcrLanguagesPl
	OcrLanguagesPt     OcrLanguages = original.OcrLanguagesPt
	OcrLanguagesRo     OcrLanguages = original.OcrLanguagesRo
	OcrLanguagesRu     OcrLanguages = original.OcrLanguagesRu
	OcrLanguagesSk     OcrLanguages = original.OcrLanguagesSk
	OcrLanguagesSrCyrl OcrLanguages = original.OcrLanguagesSrCyrl
	OcrLanguagesSrLatn OcrLanguages = original.OcrLanguagesSrLatn
	OcrLanguagesSv     OcrLanguages = original.OcrLanguagesSv
	OcrLanguagesTr     OcrLanguages = original.OcrLanguagesTr
	OcrLanguagesUnk    OcrLanguages = original.OcrLanguagesUnk
	OcrLanguagesZhHans OcrLanguages = original.OcrLanguagesZhHans
	OcrLanguagesZhHant OcrLanguages = original.OcrLanguagesZhHant
)

type TextOperationStatusCodes = original.TextOperationStatusCodes

const (
	Failed     TextOperationStatusCodes = original.Failed
	NotStarted TextOperationStatusCodes = original.NotStarted
	Running    TextOperationStatusCodes = original.Running
	Succeeded  TextOperationStatusCodes = original.Succeeded
)

type VisualFeatureTypes = original.VisualFeatureTypes

const (
	VisualFeatureTypesAdult       VisualFeatureTypes = original.VisualFeatureTypesAdult
	VisualFeatureTypesCategories  VisualFeatureTypes = original.VisualFeatureTypesCategories
	VisualFeatureTypesColor       VisualFeatureTypes = original.VisualFeatureTypesColor
	VisualFeatureTypesDescription VisualFeatureTypes = original.VisualFeatureTypesDescription
	VisualFeatureTypesFaces       VisualFeatureTypes = original.VisualFeatureTypesFaces
	VisualFeatureTypesImageType   VisualFeatureTypes = original.VisualFeatureTypesImageType
	VisualFeatureTypesTags        VisualFeatureTypes = original.VisualFeatureTypesTags
)

type AdultInfo = original.AdultInfo
type Category = original.Category
type CategoryDetail = original.CategoryDetail
type CelebritiesModel = original.CelebritiesModel
type ColorInfo = original.ColorInfo
type DomainModelResult = original.DomainModelResult
type DomainModelResults = original.DomainModelResults
type Error = original.Error
type FaceDescription = original.FaceDescription
type FaceRectangle = original.FaceRectangle
type ImageAnalysis = original.ImageAnalysis
type ImageCaption = original.ImageCaption
type ImageDescription = original.ImageDescription
type ImageDescriptionDetails = original.ImageDescriptionDetails
type ImageMetadata = original.ImageMetadata
type ImageTag = original.ImageTag
type ImageType = original.ImageType
type ImageURL = original.ImageURL
type Line = original.Line
type ListModelsResult = original.ListModelsResult
type ModelDescription = original.ModelDescription
type OcrLine = original.OcrLine
type OcrRegion = original.OcrRegion
type OcrResult = original.OcrResult
type OcrWord = original.OcrWord
type ReadCloser = original.ReadCloser
type RecognitionResult = original.RecognitionResult
type TagResult = original.TagResult
type TextOperationResult = original.TextOperationResult
type Word = original.Word

func New(azureRegion AzureRegions) BaseClient {
	return original.New(azureRegion)
}
func NewWithoutDefaults(azureRegion AzureRegions) BaseClient {
	return original.NewWithoutDefaults(azureRegion)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
