// @flow

import * as testSubject from '../generator';
import {Settings} from 'luxon';

jest.mock('uuid/v4', () => {
    return () => 'generated-uuid';
});

test('Generate access token', () => {
    Settings.now = () => 0;

    const token = testSubject.default('someUserId');

    expect(token).toEqual({
        userId: 'someUserId',
        tokenId: 'generated-uuid',
        created: new Date('1970-01-01T00:00:00.000Z'),
        expires: new Date('1971-01-01T00:00:00.000Z'),
    });
});
