// @flow

import convict from 'convict';
import manifest from './manifest.json';

export const config = convict(manifest);
config.loadFile(`./config/${config.get('env')}.json`);

config.validate({allowed: 'strict'});
