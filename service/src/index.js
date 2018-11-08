// @flow

import pino from 'pino';
import config from './config';
import buildService from './service';
import versionHandler from './service/version';
import buildDatabase from './database/db';
import buildDatabaseHealthcheck from './database/health';
import uptimeHealthchecks from './health/system/uptimeHealthcheck';
import buildHealthchecksHandler from './health/rest';
import buildAccessTokenGenerator from './authorization/generator';
import buildAccessTokenSerializer from './authorization/serializer';
import buildGenerateAccessTokenHandler from './authorization/rest/accessTokenHandler';
import buildAccessTokenParser from './authorization/rest/accessTokenParser';

const logger = pino();

const database = buildDatabase(config.get('postgres.uri'));

const accessTokenGenerator = buildAccessTokenGenerator('P1Y');
const accessTokenSerializer = buildAccessTokenSerializer('supersecretkey');

const service = buildService([
    buildAccessTokenParser(accessTokenSerializer),
], [
    versionHandler,
    buildHealthchecksHandler([uptimeHealthchecks, buildDatabaseHealthcheck(database)]),
    buildGenerateAccessTokenHandler(accessTokenSerializer, accessTokenGenerator),
]);
const port = config.get('http.port');
service.listen(port, () => logger.info({port}, 'Service started'));
