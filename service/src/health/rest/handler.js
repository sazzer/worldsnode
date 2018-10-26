// @flow

import {HEALTH_PASS, HEALTH_WARN, HEALTH_FAIL, type Health} from '../healthcheck';
import { uptimeHealthcheck } from '../system/uptimeHealthcheck';

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

/** Mapping of the health statuses to the strings to return */
const HEALTH_STATUSES = {
    [HEALTH_PASS]: 'pass',
    [HEALTH_WARN]: 'warn',
    [HEALTH_FAIL]: 'fail',
};

/**
 * Helper to compute the total status from an array of individual statuses
 */
export function computeTotalStatus(accum: string, next: string): string {
    if (accum === HEALTH_FAIL) {
        return HEALTH_FAIL;
    } else if (accum === HEALTH_WARN && next !== HEALTH_PASS) {
        return next;
    } else {
        return next;
    }
}

/**
 * Helper to build the individual component details from the otherall results
 * @param results The health results to build from
 * @return the individual components to return
 */
export function buildComponents(results: Array<Health>) : { [string]: Array<HealthCheckDetail> } {
    const details = {};
    results.forEach(result => {
        const componentName = `${result.component}:${result.measurement}`;
        const value: HealthCheckDetail = {
            componentId: result.subComponent,
            status: HEALTH_STATUSES[result.status],
            observedValue: result.value,
            observedUnit: result.unit,
            time: new Date().toISOString()
        }
        if (!details[componentName]) {
            details[componentName] = [];
        }
        details[componentName].push(value);
    });
    return details;;
}

/**
 * Actually check the health of the system
 * @param req The request
 * @param res The response
 */
export default function checkHealth(req: $Request, res: $Response) {
    const results = uptimeHealthcheck();

    const totalStatus = results.map(result => result.status)
        .reduce(computeTotalStatus, HEALTH_PASS);

    const response: HealthCheckResponse = {
        status: HEALTH_STATUSES[totalStatus],
        details: buildComponents(results)
    };

    res
        .type('application/health+json')
        .send(response);
}
