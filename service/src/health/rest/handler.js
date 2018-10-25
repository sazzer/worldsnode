// @flow

import type {
    $Request,
    $Response
} from 'express';

type HealthCheckDetail = {
    componentId?: string,
    componentType?: string,
    observedValue?: any,
    observedUnit?: string,
    status: string,
    time: string,
    output?: string
};

type HealthCheckResponse = {
    status: string,
    version?: string,
    releaseId?: string,
    notes?: Array<string>,
    output?: string,
    serviceId?: string,
    description?: string,

    details: { [string]: Array<HealthCheckDetail> }
};

/**
 * Actually check the health of the system
 * @param req The request
 * @param res The response
 */
export function checkHealth(req: $Request, res: $Response) {
    const response: HealthCheckResponse = {
        status: 'pass',
        details: {}
    };

    res
        .type('application/health+json')
        .send(response);
}
