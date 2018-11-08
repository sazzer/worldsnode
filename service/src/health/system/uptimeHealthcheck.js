// @flow

import os from 'os';
import type {Health} from '../healthcheck';
import {HEALTH_PASS} from '../healthcheck';

/**
 * Health checker that just returns the system uptime
 * @return The results of the uptime healthcheck
 */
export default function uptimeHealthcheck() : Array<Promise<Health>> {
    return [
        Promise.resolve({
            component: 'system',
            measurement: 'uptime',
            status: HEALTH_PASS,
            value: os.uptime(),
            unit: 'S',
        }),
    ];
}
