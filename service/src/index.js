// @flow

import pino from 'pino';
import config from './config';
import buildService from './service';
import versionHandler from './service/version';
import uptimeHealthchecks from './health/system/uptimeHealthcheck';
import buildHealthchecksHandler from './health/rest';

const logger = pino();

const service = buildService([
    versionHandler,
    buildHealthchecksHandler([uptimeHealthchecks]),
]);
const port = config.get('http.port');
service.listen(port, () => logger.info({port}, 'Service started'));
