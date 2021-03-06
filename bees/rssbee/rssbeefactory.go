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

package rssbee

import (
	"github.com/muesli/beehive/bees"
)

type RSSBeeFactory struct {
	bees.BeeFactory
}

// Interface impl

func (factory *RSSBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	var skip bool
	if tmp := options.GetValue("skip_first"); tmp != nil {
		skip = tmp.(bool)
	} else {
		skip = false
	}
	bee := RSSBee{
		Bee:             bees.NewBee(name, factory.Name(), description),
		url:             options.GetValue("url").(string),
		skip_next_fetch: skip,
	}

	return &bee
}

func (factory *RSSBeeFactory) Name() string {
	return "rssbee"
}

func (factory *RSSBeeFactory) Description() string {
	return "A bee that manages RSS-feeds"
}

func (factory *RSSBeeFactory) Image() string {
	return factory.Name() + ".png"
}

func (factory *RSSBeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		bees.BeeOptionDescriptor{
			Name:        "url",
			Description: "URL of the RSS-feed you want to monitor",
			Type:        "string",
			Mandatory:   true,
		},
		bees.BeeOptionDescriptor{
			Name:        "skip_first",
			Description: "Whether to skip already existing entries",
			Type:        "bool",
			Mandatory:   false,
		},
	}
	return opts
}

func (factory *RSSBeeFactory) Events() []bees.EventDescriptor {
	events := []bees.EventDescriptor{
		bees.EventDescriptor{
			Namespace:   factory.Name(),
			Name:        "newitem",
			Description: "A new item has been received through the Feed",
			Options: []bees.PlaceholderDescriptor{
				bees.PlaceholderDescriptor{
					Name:        "title",
					Description: "Title of the Item",
					Type:        "string",
				},
				bees.PlaceholderDescriptor{
					Name:        "links",
					Description: "Links referenced by the Item",
					Type:        "[]string",
				},
				bees.PlaceholderDescriptor{
					Name:        "description",
					Description: "Description of the Item",
					Type:        "string",
				},
				bees.PlaceholderDescriptor{
					Name:        "author",
					Description: "The person who wrote the Item",
					Type:        "string",
				},
				bees.PlaceholderDescriptor{
					Name:        "categories",
					Description: "Categories that the Item belongs to",
					Type:        "[]string",
				},
				bees.PlaceholderDescriptor{
					Name:        "comments",
					Description: "Comments of the Item",
					Type:        "string",
				},
				bees.PlaceholderDescriptor{
					Name:        "enclosures",
					Description: "Enclosures related to Item",
					Type:        "[]string",
				},
				bees.PlaceholderDescriptor{
					Name:        "guid",
					Description: "Global unique ID attached to the Item",
					Type:        "string",
				},
				bees.PlaceholderDescriptor{
					Name:        "pubdate",
					Description: "Date the Item was published on",
					Type:        "string",
				},
				bees.PlaceholderDescriptor{
					Name:        "source",
					Description: "Source of the Item",
					Type:        "string",
				},
			},
		},
	}
	return events
}

func init() {
	f := RSSBeeFactory{}
	bees.RegisterFactory(&f)
}
