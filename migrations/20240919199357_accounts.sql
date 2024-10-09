-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    email VARCHAR(100),
    password VARCHAR(100),
    number SERIAL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    balance INTEGER
);
INSERT INTO accounts (username, email, password, balance) VALUES
    ('Grimes', 'astral.wanderer@neonrealm.com', 'password123', 16326326),
    ('MatthewBellamy', 'grimes@gmail.com', 'password123', 503623),
    ('EchoDancer', 'echo.dancer@lunalands.net', 'password123', 1262363),
    ('Crystaella', 'crystaella@elvenverse.org', 'password123', 75141),
    ('SkyGlider', 'sky.glider@archeageworld.com', 'password123', 98454),
    ('NovaRune', 'nova.rune@celestialwinds.net', 'password123', 68677),
    ('NebulaDrifter', 'nebula.drifter@cosmicsymphony.org', 'password123', 85796979),
    ('EtherBlossom', 'ether.blossom@solflower.com', 'password123', 197696),
    ('QuantumTide', 'quantum.tide@tidebreakers.com', 'password123', 74766),
    ('LunaPulse', 'luna.pulse@twilightsea.org', 'password123', 6842),
    ('AetherKnight', 'aether.knight@stormrealms.com', 'password123', 9505487),
    ('Velorion', 'velorion@ancientwoods.com', 'password123', 778654),
    ('StellarWhisper', 'stellar.whisper@starsong.net', 'password123', 89825),
    ('Glintshade', 'glintshade@eclipsehollow.org', 'password123', 53825),
    ('AzureMist', 'azure.mist@oceansoul.net', 'password123', 46825),
    ('Dreamweaver', 'dreamweaver@chronobreakers.com', 'password123', 64825),
    ('SolarFlare', 'solar.flare@firespirits.net', 'password123', 97825),
    ('TwilightVoyager', 'twilight.voyager@skiesend.org', 'password123', 83825),
    ('LunarEcho', 'lunar.echo@moonlitrealm.com', 'password123', 51825),
    ('PhantomFlicker', 'phantom.flicker@veilshadows.com', 'password123', 72825),
    ('CosmoWisp', 'cosmo.wisp@astraljourney.org', 'password123', 105825),
    ('VoidEcho', 'void.echo@neonpath.com', 'password123', 68825),
    ('StardustGleam', 'stardust.gleam@cosmicwave.com', 'password123', 87825),
    ('SolarWind', 'solar.wind@arcanewinds.net', 'password123', 62825),
    ('EtherealDawn', 'ethereal.dawn@luminoushorizon.com', 'password123', 54825),
    ('GalacticShade', 'galactic.shade@stardewrise.org', 'password123', 92825),
    ('CrescentSong', 'crescent.song@twinmoonrealms.net', 'password123', 76825);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE accounts;
-- +goose StatementEnd
