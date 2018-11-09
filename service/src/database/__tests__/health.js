// @flow

import * as testSubject from '../health';

describe('Database Healthcheck', () => {
    test('Returns Success', () => {
        const db = {
            query: jest.fn(() => Promise.resolve([{'column': 1}])),
        };

        const result = testSubject.databaseHealthcheck(db);
        expect(result).toHaveLength(1);
        expect(result[0]).resolves.toEqual({
            component: 'database',
            measurement: 'ping',
            status: 'HEALTH_PASS',
        });
        expect(db.query).toBeCalledWith({sql: 'SELECT 1'});
    });
    test('Returns Failure', () => {
        const db = {
            query: jest.fn(() => Promise.reject()),
        };

        const result = testSubject.databaseHealthcheck(db);
        expect(result).toHaveLength(1);
        expect(result[0]).resolves.toEqual({
            component: 'database',
            measurement: 'ping',
            status: 'HEALTH_FAIL',
        });
        expect(db.query).toBeCalledWith({sql: 'SELECT 1'});
    });
});
