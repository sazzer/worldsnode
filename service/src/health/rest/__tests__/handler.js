// @flow

import {Response} from 'jest-express/lib/response';
import * as testSubject from '../handler';

describe('computeTotalStatus', () => {
    test.each`
        before           | next             | expected
        ${'HEALTH_PASS'} | ${'HEALTH_PASS'} | ${'HEALTH_PASS'}
        ${'HEALTH_PASS'} | ${'HEALTH_WARN'} | ${'HEALTH_WARN'}
        ${'HEALTH_PASS'} | ${'HEALTH_FAIL'} | ${'HEALTH_FAIL'}
        ${'HEALTH_WARN'} | ${'HEALTH_PASS'} | ${'HEALTH_WARN'}
        ${'HEALTH_WARN'} | ${'HEALTH_WARN'} | ${'HEALTH_WARN'}
        ${'HEALTH_WARN'} | ${'HEALTH_FAIL'} | ${'HEALTH_FAIL'}
        ${'HEALTH_FAIL'} | ${'HEALTH_PASS'} | ${'HEALTH_FAIL'}
        ${'HEALTH_FAIL'} | ${'HEALTH_WARN'} | ${'HEALTH_FAIL'}
        ${'HEALTH_FAIL'} | ${'HEALTH_FAIL'} | ${'HEALTH_FAIL'}
        `('returns $expected when computing $next from $before',
    ({before, next, expected}) => {
        const output = testSubject.computeTotalStatus(before, next);
        expect(output).toEqual(expected);
    });
});

describe('buildComponents', () => {
    test('No health checks', () => {
        const result = testSubject.buildComponents([]);
        expect(result).toEqual({});
    });
    test('One health checks', () => {
        const uptime = {
            component: 'system',
            measurement: 'uptime',
            status: 'HEALTH_PASS',
            value: 5,
            unit: 'S',
        };
        const result = testSubject.buildComponents([uptime]);
        expect(result).toEqual({
            'system:uptime': [{
                'componentId': undefined,
                'status': 'pass',
                'observedValue': 5,
                'observedUnit': 'S',
            }],
        });
    });
    test('Duplicate health checks', () => {
        const uptime1 = {
            component: 'system',
            measurement: 'uptime',
            status: 'HEALTH_PASS',
            value: 5,
            unit: 'S',
        };
        const uptime2 = {
            component: 'system',
            measurement: 'uptime',
            status: 'HEALTH_PASS',
            value: 15,
            unit: 'S',
        };
        const result = testSubject.buildComponents([uptime1, uptime2]);
        expect(result).toEqual({
            'system:uptime': [{
                'componentId': undefined,
                'status': 'pass',
                'observedValue': 5,
                'observedUnit': 'S',
            }, {
                'componentId': undefined,
                'status': 'pass',
                'observedValue': 15,
                'observedUnit': 'S',
            }],
        });
    });
    test('Different health checks', () => {
        const uptime = {
            component: 'system',
            measurement: 'uptime',
            status: 'HEALTH_PASS',
            value: 5,
            unit: 'S',
        };
        const disk = {
            component: 'system',
            measurement: 'disk',
            status: 'HEALTH_PASS',
            value: 15,
            unit: 'MiB',
        };
        const result = testSubject.buildComponents([uptime, disk]);
        expect(result).toEqual({
            'system:uptime': [{
                'componentId': undefined,
                'status': 'pass',
                'observedValue': 5,
                'observedUnit': 'S',
            }],
            'system:disk': [{
                'componentId': undefined,
                'status': 'pass',
                'observedValue': 15,
                'observedUnit': 'MiB',
            }],
        });
    });
});

describe('checkHealth', () => {
    test('No healthchecks', () => {
        const response = new Response();
        testSubject.default(response, []);

        expect(response.type).toBeCalledWith('application/health+json');
        expect(response.status).toBeCalledWith(200);
        expect(response.send).toBeCalledWith({
            status: 'pass',
            details: {},
        });
    });
    test('One passing healthcheck', () => {
        const healthcheck = jest.fn();
        healthcheck.mockReturnValue([{
            component: 'system',
            measurement: 'uptime',
            status: 'HEALTH_PASS',
            value: 5,
            unit: 'S',
        }]);

        const response = new Response();
        testSubject.default(response, [healthcheck]);

        expect(response.type).toBeCalledWith('application/health+json');
        expect(response.status).toBeCalledWith(200);
        expect(response.send).toBeCalledWith({
            status: 'pass',
            details: {
                'system:uptime': [{
                    'componentId': undefined,
                    'observedUnit': 'S',
                    'observedValue': 5,
                    'status': 'pass',
                }],
            },
        });
    });
    test('One warning healthcheck', () => {
        const healthcheck = jest.fn();
        healthcheck.mockReturnValue([{
            component: 'system',
            measurement: 'uptime',
            status: 'HEALTH_WARN',
            value: 5,
            unit: 'S',
        }]);

        const response = new Response();
        testSubject.default(response, [healthcheck]);

        expect(response.type).toBeCalledWith('application/health+json');
        expect(response.status).toBeCalledWith(200);
        expect(response.send).toBeCalledWith({
            status: 'warn',
            details: {
                'system:uptime': [{
                    'componentId': undefined,
                    'observedUnit': 'S',
                    'observedValue': 5,
                    'status': 'warn',
                }],
            },
        });
    });
    test('One failing healthcheck', () => {
        const healthcheck = jest.fn();
        healthcheck.mockReturnValue([{
            component: 'system',
            measurement: 'uptime',
            status: 'HEALTH_FAIL',
            value: 5,
            unit: 'S',
        }]);

        const response = new Response();
        testSubject.default(response, [healthcheck]);

        expect(response.type).toBeCalledWith('application/health+json');
        expect(response.status).toBeCalledWith(500);
        expect(response.send).toBeCalledWith({
            status: 'fail',
            details: {
                'system:uptime': [{
                    'componentId': undefined,
                    'observedUnit': 'S',
                    'observedValue': 5,
                    'status': 'fail',
                }],
            },
        });
    });
});
