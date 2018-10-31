// @flow

import {buildAccessToken} from '../accessTokenHandler';
import {Request, Response} from 'jest-express';

describe('buildAccessToken', () => {
    test('Successfully provide an access token', async () => {
        const accessToken = {
            userId: 'someUserId',
            tokenId: 'generated-uuid',
            created: new Date('1970-01-01T00:00:00.000Z'),
            expires: new Date('1971-01-01T00:00:00.000Z'),
        };

        const generator = {
            generate: jest.fn(() => accessToken),
        };
        const serializer = {
            serialize: jest.fn(() => Promise.resolve('generatedJwt')),
        };
        const request = new Request('/api/debug/accessToken');
        request.setBody({userId: 'someUserId'});
        const response = new Response();

        await buildAccessToken(request, response, serializer, generator);

        expect(generator.generate).toBeCalledWith('someUserId');
        expect(serializer.serialize).toBeCalledWith(accessToken);
        expect(response.send).toBeCalledWith({
            token: 'generatedJwt',
        });
    });
});
