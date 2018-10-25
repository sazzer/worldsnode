// @flow

import convict from 'convict';
import manifest from './manifest.json';

export const config = convict(manifest);
config.validate({allowed: 'strict'});
