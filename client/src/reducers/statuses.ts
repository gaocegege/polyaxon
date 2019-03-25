import { normalize } from 'normalizr';
import { Reducer } from 'redux';

import * as _ from 'lodash';

import { actionTypes, StatusesAction } from '../actions/statuses';
import { ACTIONS } from '../constants/actions';
import { StatusSchema } from '../constants/schemas';
import { ErrorEmptyState, ErrorSchema, processErrorGlobal } from '../models/errors';
import {
  LoadingIndicatorEmptyState,
  LoadingIndicatorSchema,
  processLoadingIndicatorGlobal
} from '../models/loadingIndicator';
import { StatusEmptyState, StatusModel, StatusStateSchema } from '../models/status';
import { LastFetchedIds } from '../models/utils';

export const StatusesReducer: Reducer<StatusStateSchema> =
  (state: StatusStateSchema = StatusEmptyState, action: StatusesAction) => {
    let newState = {...state};

    const processStatus = function (status: StatusModel) {
      const id = status.id;
      newState.lastFetched.ids.push(id);
      if (!_.includes(newState.ids, id)) {
        newState.ids.push(id);
      }
      const normalizedStatuses = normalize(status, StatusSchema).entities.statuses;
      newState.byIds[id] = {
        ...newState.byIds[id], ...normalizedStatuses[id]
      };
      return newState;
    };

    switch (action.type) {
      case actionTypes.FETCH_STATUSES_REQUEST:
        newState.lastFetched = new LastFetchedIds();
        return newState;
      case actionTypes.FETCH_STATUSES_SUCCESS:
        newState.lastFetched = new LastFetchedIds();
        newState.lastFetched.count = action.count;
        for (const build of action.statuses) {
          newState = processStatus(build);
        }
        return newState;
      default:
        return state;
    }
  };

export const LoadingIndicatorStatusesReducer: Reducer<LoadingIndicatorSchema> =
  (state: LoadingIndicatorSchema = LoadingIndicatorEmptyState, action: StatusesAction) => {
    switch (action.type) {
      case actionTypes.FETCH_STATUSES_REQUEST:
        return {
          ...state,
          statuses: processLoadingIndicatorGlobal(state.statuses, true, ACTIONS.FETCH)
        };
      case actionTypes.FETCH_STATUSES_ERROR:
      case actionTypes.FETCH_STATUSES_SUCCESS:
        return {
          ...state,
          statuses: processLoadingIndicatorGlobal(state.statuses, false, ACTIONS.FETCH)
        };
      default:
        return state;
    }
  };

export const ErrorStatusesReducer: Reducer<ErrorSchema> =
  (state: ErrorSchema = ErrorEmptyState, action: StatusesAction) => {
    switch (action.type) {
      case actionTypes.FETCH_STATUSES_REQUEST:
      case actionTypes.FETCH_STATUSES_SUCCESS:
        return {
          ...state,
          statuses: processErrorGlobal(state.statuses, null, ACTIONS.FETCH)
        };
      case actionTypes.FETCH_STATUSES_ERROR:
        return {
          ...state,
          statuses: processErrorGlobal(state.statuses, action.error, ACTIONS.FETCH)
        };
      default:
        return state;
    }
  };
