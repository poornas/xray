/*
 * Copyright (c) 2017 Minio, Inc. <https://www.minio.io>
 *
 * This file is part of Xray.
 *
 * Xray is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package cmd

import "github.com/lazywei/go-opencv/opencv"

// Global constants for Xray.
const (
	minGoVersion = ">= 1.7" // Xray requires at least Go v1.7
)

var (
	// Default haar cascade classifier used for detecting faces.
	globalHaarCascadeClassifier = opencv.LoadHaarClassifierCascade("haarcascade_frontalface_alt.xml")

	globalXrayCertFile = "/etc/ssl/public.crt"
	globalXrayKeyFile  = "/etc/ssl/private.key"

	// List of all other classifiers.
)
