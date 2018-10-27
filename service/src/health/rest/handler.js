// @flow

import {
    HEALTH_PASS,
    HEALTH_WARN,
    HEALTH_FAIL,
    type Health,
    type HealthChecker,
} from '../healthcheck';

import {
    type $Response,
} from 'express';

type HealthCheckDetail = {
    componentId?: string,
    observedValue?: any,
    observedUnit?: string,
    status: string,
};

type HealthCheckResponse = {
    status: string,
    details: { [string]: Array<HealthCheckDetail> }
};

/** Mapping of the health statuses to the strings to return */
const HEALTH_STATUSES = {
    [HEALTH_PASS]: 'pass',
    [HEALTH_WARN]: 'warn',
    [HEALTH_FAIL]: 'fail',
};

/** Mapping of the health statuses to the HTTP Status Codes to return */
const HEALTH_STATUS_CODES = {
    [HEALTH_PASS]: 200,
    [HEALTH_WARN]: 200,
    [HEALTH_FAIL]: 500,
};

/**
 * Helper to compute the total status from an array of individual statuses
 * @param accum The accumulated value so far
 * @param next The next value to compute from
 * @return the newly computed value
 */
export function computeTotalStatus(accum: string, next: string): string {
    if (accum === HEALTH_FAIL || next === HEALTH_FAIL) {
        return HEALTH_FAIL;
    } else if (accum === HEALTH_WARN) {
        return HEALTH_WARN;
    } else {
        return next;
    }
}

/**
 * Helper to build the individual component details from the otherall results
 * @param results The health results to build from
 * @return the individual components to return
 */
export function buildComponents(results: Array<Health>): { [string]: Array<HealthCheckDetail> } {
    const details = {};
    results.forEach((result) => {
        const componentName = `${result.component}:${result.measurement}`;
        const value: HealthCheckDetail = {
            componentId: result.subComponent,
            status: HEALTH_STATUSES[result.status],
            observedValue: result.value,
            observedUnit: result.unit,
        };
        if (!details[componentName]) {
            details[componentName] = [];
        }
        details[componentName].push(value);
    });
    return details;
}

/**
 * Actually check the health of the system
 * @param res The response to write to
 * @param healthchecks The healthchecks to perform
 */
export default function checkHealth(res: $Response, healthchecks: Array<HealthChecker>) {
    const results = healthchecks
        .map((healthcheck) => healthcheck())
        .reduce((a, b) => a.concat(b), []);

    const totalStatus = results.map((result) => result.status)
        .reduce(computeTotalStatus, HEALTH_PASS);

    const response: HealthCheckResponse = {
        status: HEALTH_STATUSES[totalStatus],
        details: buildComponents(results),
    };

    res
        .type('application/health+json')
        .status(HEALTH_STATUS_CODES[totalStatus])
        .send(response);
}
