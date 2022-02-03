// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pdusession

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReleasePduSession - Release
func ReleasePduSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdatePduSession - Update (initiated by V-SMF)
func UpdatePduSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
