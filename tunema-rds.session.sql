-- INSERT INTO tags(category_id, name)
--     VALUES (1, 'Hip Hop'),
-- (1, 'House / Techno'),
-- (1, 'Pop / EDM'),
-- (1, 'Live Sounds'),
-- (1, 'Electronic'),
-- (2, 'Drums'),
-- (2, 'Kicks'),
-- (2, 'Snares'),
-- (2, 'Hats'),
-- (2, 'Claps'),
-- (2, 'Toms'),
-- (2, 'Cymbals'),
-- (2, '808'),
-- (3, 'Female'),
-- (3, 'Male'),
-- (3, 'FX'),
-- (3, 'Spoken Word'),
-- (3, 'Phrases'),
-- (3, 'Screams'),
-- (3, 'Shouts');
-- INSERT INTO samples(user_id, name, bpm, key, key_scale, time, file_url, cover_url, price, created_at)
--     VALUES (1, 'Gh0st_wave_Surround_Sound_125BPM_Am.wav', 125, 'A', 'Minor', 10, '', '', 19.99, now()),
-- (2, 'KAELIN_ELLIS_90_drum_loop_chestpuncher.wav', 90, '', '', 1, '', '', 1.99, now());
-- SELECT
--     samples.*,
--     users.username AS author_name
-- FROM
--     samples
--     LEFT JOIN users ON samples.user_id = users.id;
-- insert into sample_tags (sample_id, tag_id)
-- values (1, 1),
--         (2, 1),
--         (2, 6),
--         (2, 7),
--         (2, 8),
--         (2, 9);
-- SELECT
--     samples.*,
--     users.username AS author_name,
--     ARRAY_AGG(tags.name) AS tags,
--     COUNT(order_products.sample_id) AS sold
-- FROM
--     samples
--     LEFT JOIN users ON samples.user_id = users.id
--     LEFT JOIN sample_tags ON samples.id = sample_tags.sample_id
--     LEFT JOIN tags ON sample_tags.tag_id = tags.id
--     LEFT JOIN order_products ON order_products.sample_id = samples.id
-- GROUP BY
--     samples.id,
--     users.username;
-- SELECT
--     tags.id AS tag_id,
--     tags.name AS tag_name,
--     categories.name AS category_name
-- FROM
--     tags
--     LEFT JOIN categories ON tags.category_id = categories.id;
-- SELECT
--     samples.*,
--     users.username AS artist_name,
--     ARRAY_AGG(tags.name) AS tags,
--     COUNT(order_products.sample_id) AS sold
-- FROM
--     samples
--     LEFT JOIN users ON samples.user_id = users.id
--     LEFT JOIN sample_tags ON samples.id = sample_tags.sample_id
--     LEFT JOIN tags ON sample_tags.tag_id = tags.id
--     LEFT JOIN order_products ON order_products.sample_id = samples.id
-- GROUP BY
--     samples.id,
--     users.username
-- HAVING
--     ARRAY_AGG(tags.id) @> ARRAY[16]::integer[]
-- SELECT
--     samples.*
-- FROM
--     samples
--     LEFT JOIN users ON samples.user_id = users.id
--     LEFT JOIN sample_tags ON samples.id = sample_tags.sample_id
--     LEFT JOIN tags ON sample_tags.tag_id = tags.id
--     LEFT JOIN order_products ON order_products.sample_id = samples.id
-- WHERE
--     samples.user_id = 1
-- SELECT
--     samples.*,
--     users.username AS artist_name,
--     ARRAY_AGG(tags.name) AS tags,
--     COUNT(order_products.sample_id) AS sold
-- FROM
--     samples
--     LEFT JOIN users ON samples.user_id = users.id
--     LEFT JOIN sample_tags ON samples.id = sample_tags.sample_id
--     LEFT JOIN tags ON sample_tags.tag_id = tags.id
--     LEFT JOIN order_products ON order_products.sample_id = samples.id
-- WHERE
--     samples.user_id = 2
-- GROUP BY
--     samples.id,
--     users.username
-- ORDER BY
--     created_at DESC
-- UPDATE
--     users
-- SET
--     profile_img_url = 'https://hostedboringavatars.vercel.app/api/beam?colors=EAF89B,141414&name=' || username
-- WHERE
--     profile_img_url = 'https://i.pravatar.cc/300'
SELECT
    *
FROM
    users;

