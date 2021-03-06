// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.0.90
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    V1Version,
    V1VersionFromJSON,
    V1VersionFromJSONTyped,
    V1VersionToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1Versions
 */
export interface V1Versions {
    /**
     * 
     * @type {string}
     * @memberof V1Versions
     */
    platform_version?: string;
    /**
     * 
     * @type {V1Version}
     * @memberof V1Versions
     */
    cli?: V1Version;
    /**
     * 
     * @type {V1Version}
     * @memberof V1Versions
     */
    platform?: V1Version;
    /**
     * 
     * @type {V1Version}
     * @memberof V1Versions
     */
    agent?: V1Version;
}

export function V1VersionsFromJSON(json: any): V1Versions {
    return V1VersionsFromJSONTyped(json, false);
}

export function V1VersionsFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Versions {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'platform_version': !exists(json, 'platform_version') ? undefined : json['platform_version'],
        'cli': !exists(json, 'cli') ? undefined : V1VersionFromJSON(json['cli']),
        'platform': !exists(json, 'platform') ? undefined : V1VersionFromJSON(json['platform']),
        'agent': !exists(json, 'agent') ? undefined : V1VersionFromJSON(json['agent']),
    };
}

export function V1VersionsToJSON(value?: V1Versions | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'platform_version': value.platform_version,
        'cli': V1VersionToJSON(value.cli),
        'platform': V1VersionToJSON(value.platform),
        'agent': V1VersionToJSON(value.agent),
    };
}


