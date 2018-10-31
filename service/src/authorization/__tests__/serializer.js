// @flow

import * as testSubject from '../serializer';

/** The signing key */
const signingKey = 'somesupersecretkey';

/** The token being serialized */
const token = {
    userId: 'someUserId',
    tokenId: 'generated-uuid',
    created: new Date('2000-01-01T00:00:00.000Z'),
    expires: new Date('2170-01-01T00:00:00.000Z'),
};

/** The serialized form of the above token */
/* eslint-disable max-len */
const serialized = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ3b3JsZHMiLCJpc3MiOiJ3b3JsZHMiLCJqd3RpZCI6ImdlbmVyYXRlZC11dWlkIiwic3ViIjoic29tZVVzZXJJZCIsImlhdCI6OTQ2Njg0ODAwLCJleHAiOjYzMTE0MzM2MDB9.UtKVmCIMSiZrw2ynr6vV7JJEkLwRMhw3Vo8QJsyOVM4';
/* eslint-enable max-len */

describe('Exported functions', () => {
    test('Serialize an access token', async () => {
        const result = await testSubject.serialize(token, signingKey);
        expect(result).toEqual(serialized);
    });

    test('Deserialize an access token', async () => {
        const result = await testSubject.deserialize(serialized, signingKey);
        expect(result).toEqual(token);
    });

    test('Deserialize an invalid access token', () => {
        return expect(testSubject.deserialize('invalid', signingKey)).rejects.toEqual({
            message: 'jwt malformed',
        });
    });
});

describe('Built Serializer', () => {
    const serializer = testSubject.default(signingKey);

    test('Serialize an access token', async () => {
        const result = await serializer.serialize(token);
        expect(result).toEqual(serialized);
    });

    test('Deserialize an access token', async () => {
        const result = await serializer.deserialize(serialized);
        expect(result).toEqual(token);
    });

    test('Deserialize an invalid access token', () => {
        return expect(serializer.deserialize('invalid')).rejects.toEqual({
            message: 'jwt malformed',
        });
    });
});
