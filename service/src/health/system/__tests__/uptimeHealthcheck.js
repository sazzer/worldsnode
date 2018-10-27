// @flow

import uptimeHealthcheck from '../uptimeHealthcheck';
import os from 'os';

test('Uptime Healthcheck', () => {
    os.uptime = jest.fn();
    os.uptime.mockReturnValueOnce(12345);

    const result = uptimeHealthcheck();
    expect(result).toHaveLength(1);
    expect(result[0]).toEqual({
        component: 'system',
        measurement: 'uptime',
        status: 'HEALTH_PASS',
        value: 12345,
        unit: 'S',
    });
});
