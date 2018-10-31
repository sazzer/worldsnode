// @flow

import * as testSubject from '../generator';
import {Duration, Settings} from 'luxon';

jest.mock('uuid/v4', () => {
    return () => 'generated-uuid';
});

describe('Exported functions', () => {
    test('Generate access token', () => {
        Settings.now = () => 0;

        const token = testSubject.generate('someUserId', Duration.fromISO('P1Y'));

        expect(token).toEqual({
            userId: 'someUserId',
            tokenId: 'generated-uuid',
            created: new Date('1970-01-01T00:00:00.000Z'),
            expires: new Date('1971-01-01T00:00:00.000Z'),
        });
    });
});

describe('Builder', () => {
    test('Generate access token', () => {
        const generator = testSubject.default('P1Y');
        Settings.now = () => 0;

        const token = generator.generate('someUserId');

        expect(token).toEqual({
            userId: 'someUserId',
            tokenId: 'generated-uuid',
            created: new Date('1970-01-01T00:00:00.000Z'),
            expires: new Date('1971-01-01T00:00:00.000Z'),
        });
    });
});
