import {pluginId, pluginVersion} from './manifest';

// To ease migration, verify separate export of id and version.
test('Plugin id and version are defined', () => {
    expect(pluginId).toBeDefined();
    expect(pluginVersion).toBeDefined();
});
