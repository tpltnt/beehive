/*
 *    Copyright (C) 2014 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

package huebee

import (
	"github.com/muesli/beehive/bees"
)

type HueBeeFactory struct {
	bees.BeeFactory
}

// Interface impl

func (factory *HueBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := HueBee{
		Bee: bees.NewBee(name, factory.Name(), description),
		bridge: options.GetValue("bridge").(string),
		key: options.GetValue("key").(string),
	}

	return &bee
}

func (factory *HueBeeFactory) Name() string {
	return "huebee"
}

func (factory *HueBeeFactory) Description() string {
	return "A Philips hue module for beehive"
}

func (factory *HueBeeFactory) Image() string {
	return factory.Name() + ".png"
}

func (factory *HueBeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		bees.BeeOptionDescriptor{
			Name:        "bridge",
			Description: "Address of the Hue bridge, eg: 192.168.0.1",
			Type:        "url",
			Mandatory:   true,
		},
		bees.BeeOptionDescriptor{
			Name:        "key",
			Description: "Key used for auth with the bridge",
			Type:        "string",
			Mandatory:   true,
		},
	}
	return opts
}

func (factory *HueBeeFactory) Events() []bees.EventDescriptor {
	events := []bees.EventDescriptor{}
	return events
}

func (factory *HueBeeFactory) Actions() []bees.ActionDescriptor {
	actions := []bees.ActionDescriptor{
		bees.ActionDescriptor{
			Namespace:   factory.Name(),
			Name:        "switch",
			Description: "Switches on/off a Hue light",
			Options: []bees.PlaceholderDescriptor{
				bees.PlaceholderDescriptor{
					Name:        "light",
					Description: "ID of the light you want to switch on or off",
					Type:        "int",
				},
				bees.PlaceholderDescriptor{
					Name:        "state",
					Description: "New state of the light, true for turning it on",
					Type:        "bool",
				},
			},
		},
	}
	return actions
}

func init() {
	f := HueBeeFactory{}
	bees.RegisterFactory(&f)
}
