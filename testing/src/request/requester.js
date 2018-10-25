// @flow

import { config } from '../config';
import axios from 'axios';

/**
 * The shape of the request details
 */
export type Request = {
    method?: string
};

/**
 * The shape of the response from a request
 */
export type Response = {
    status: number,
    data: any
};

/**
 * Actually make a request to the service
 * @param url the URL to request
 * @param request the Request details
 */
export function request(url: string, request: Request = {}): Promise<Response> {
    return axios.request({
        baseURL: config.get('service.uri'),
        url: url,
        timeout: 5000,
        method: request.method || 'GET',
        validateStatus: (status: number) => true
    });
}
