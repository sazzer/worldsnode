// @flow

/** Type representing the actual health status of some component of the system */
export type Health = {
    component: string,
    measurement: string,
    status: string,
    subComponent?: string,
    value?: any,
    unit?: string
};

/** Value representing a passing healthcheck */
export const HEALTH_PASS = 'HEALTH_PASS';
/** Value representing a warning healthcheck */
export const HEALTH_WARN = 'HEALTH_WARN';
/** Value representing a failing healthcheck */
export const HEALTH_FAIL = 'HEALTH_FAIL';

/** Type representing something that can check some health components */
export type HealthChecker = () => Array<Promise<Health>>;
