/*
 * Copyright 2019 Armory, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package plank

import "fmt"

func (notifications *NotificationsType) FillAppNotificationFields(appName string) {
	notificationsMap := *notifications
	for notApp, sliceOfNotifications := range notificationsMap{
		if mapped, ok := sliceOfNotifications.(map[string]interface{}); ok {
			//for _, currentNotification := range mapped {
			//currentValues := currentNotification.(map[string]interface{})
			mapped["level"] = "application"
			mapped["type"] = notApp
		}
	}
	notificationsMap["application"] = appName
}


// GetApplicationNotifications returns all application notifications
func (c *Client) GetApplicationNotifications(appName string) (*NotificationsType, error) {
	var notifications NotificationsType
	if err := c.Get(c.URLs["front50"]+"/notifications/application/"+appName, &notifications); err != nil {
		return nil, err
	}
	return &notifications, nil
}

// UpdateApplicationNotifications updates notifications in the configured front50 store.
func (c *Client) UpdateApplicationNotifications(notifications NotificationsType, appName string) error {
	notifications.FillAppNotificationFields(appName)
	var unused interface{}
	if err := c.Post(fmt.Sprintf("%s/notifications/application/%s", c.URLs["front50"], appName), ApplicationJson, notifications , &unused); err != nil {
		return fmt.Errorf("could not update notifications %q: %w", notifications, err)
	}
	return nil
}