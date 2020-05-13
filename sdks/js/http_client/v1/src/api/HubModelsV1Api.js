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

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.0.89
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import RuntimeError from '../model/RuntimeError';
import V1HubModel from '../model/V1HubModel';
import V1ListHubModelsResponse from '../model/V1ListHubModelsResponse';

/**
* HubModelsV1 service.
* @module api/HubModelsV1Api
* @version 1.0.89
*/
export default class HubModelsV1Api {

    /**
    * Constructs a new HubModelsV1Api. 
    * Polyaxon sdk
    * @alias module:api/HubModelsV1Api
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the createHubModel operation.
     * @callback module:api/HubModelsV1Api~createHubModelCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HubModel} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create hub model
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1HubModel} body Model body
     * @param {module:api/HubModelsV1Api~createHubModelCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HubModel}
     */
    createHubModel(owner, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling createHubModel");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createHubModel");
      }

      let pathParams = {
        'owner': owner
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1HubModel;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteHubModel operation.
     * @callback module:api/HubModelsV1Api~deleteHubModelCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete hub model
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/HubModelsV1Api~deleteHubModelCallback} callback The callback function, accepting three arguments: error, data, response
     */
    deleteHubModel(owner, uuid, callback) {
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling deleteHubModel");
      }
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling deleteHubModel");
      }

      let pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{uuid}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getHubModel operation.
     * @callback module:api/HubModelsV1Api~getHubModelCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HubModel} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get hub model
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/HubModelsV1Api~getHubModelCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HubModel}
     */
    getHubModel(owner, uuid, callback) {
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getHubModel");
      }
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling getHubModel");
      }

      let pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1HubModel;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{uuid}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listHubModelNames operation.
     * @callback module:api/HubModelsV1Api~listHubModelNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListHubModelsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List hub model names
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/HubModelsV1Api~listHubModelNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListHubModelsResponse}
     */
    listHubModelNames(owner, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listHubModelNames");
      }

      let pathParams = {
        'owner': owner
      };
      let queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1ListHubModelsResponse;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/names', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listHubModels operation.
     * @callback module:api/HubModelsV1Api~listHubModelsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListHubModelsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List hub models
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/HubModelsV1Api~listHubModelsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListHubModelsResponse}
     */
    listHubModels(owner, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listHubModels");
      }

      let pathParams = {
        'owner': owner
      };
      let queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1ListHubModelsResponse;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the patchHubModel operation.
     * @callback module:api/HubModelsV1Api~patchHubModelCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HubModel} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch hub model
     * @param {String} owner Owner of the namespace
     * @param {String} model_uuid UUID
     * @param {module:model/V1HubModel} body Model body
     * @param {module:api/HubModelsV1Api~patchHubModelCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HubModel}
     */
    patchHubModel(owner, model_uuid, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling patchHubModel");
      }
      // verify the required parameter 'model_uuid' is set
      if (model_uuid === undefined || model_uuid === null) {
        throw new Error("Missing the required parameter 'model_uuid' when calling patchHubModel");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling patchHubModel");
      }

      let pathParams = {
        'owner': owner,
        'model.uuid': model_uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1HubModel;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{model.uuid}', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the updateHubModel operation.
     * @callback module:api/HubModelsV1Api~updateHubModelCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HubModel} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update hub model
     * @param {String} owner Owner of the namespace
     * @param {String} model_uuid UUID
     * @param {module:model/V1HubModel} body Model body
     * @param {module:api/HubModelsV1Api~updateHubModelCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HubModel}
     */
    updateHubModel(owner, model_uuid, body, callback) {
      let postBody = body;
      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateHubModel");
      }
      // verify the required parameter 'model_uuid' is set
      if (model_uuid === undefined || model_uuid === null) {
        throw new Error("Missing the required parameter 'model_uuid' when calling updateHubModel");
      }
      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateHubModel");
      }

      let pathParams = {
        'owner': owner,
        'model.uuid': model_uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = V1HubModel;
      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/models/{model.uuid}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
