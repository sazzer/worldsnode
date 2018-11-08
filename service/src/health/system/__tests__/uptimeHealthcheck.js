// @flow

import uptimeHealthcheck from '../uptimeHealthcheck';

jest.mock('os', () => {
    return {
        uptime: () => 12345,
    };
});
test('Uptime Healthcheck', () => {
    const result = uptimeHealthcheck();
    expect(result).toHaveLength(1);
    expect(result[0]).resolves.toEqual({
        component: 'system',
        measurement: 'uptime',
        status: 'HEALTH_PASS',
        value: 12345,
        unit: 'S',
    });
});
