INSERT INTO theaters ("name", "location",country, rooms)
VALUES
('Cineplex Paradise', '123 Main Street','USA', 6),
('MegaPlex Cinemas', '456 Elm Street','England', 5);

INSERT INTO rooms (room_number, seats, "rows", theater_id)
VALUES
(1, 100, 10, (SELECT theater_id FROM theaters WHERE "name" = 'Cineplex Paradise')),
(2, 80, 8, (SELECT theater_id FROM theaters WHERE "name" = 'Cineplex Paradise')),
(3, 70, 7, (SELECT theater_id FROM theaters WHERE "name" = 'Cineplex Paradise')),
(4, 60, 6, (SELECT theater_id FROM theaters WHERE "name" = 'Cineplex Paradise')),
(5, 50, 5, (SELECT theater_id FROM theaters WHERE "name" = 'Cineplex Paradise'));

-- Insert rooms for 'MegaPlex Cinemas'
INSERT INTO rooms (room_number, seats, "rows", theater_id)
VALUES
(1, 100, 10, (SELECT theater_id FROM theaters WHERE "name" = 'MegaPlex Cinemas')),
(2, 90, 9, (SELECT theater_id FROM theaters WHERE "name" = 'MegaPlex Cinemas')),
(3, 80, 8, (SELECT theater_id FROM theaters WHERE "name" = 'MegaPlex Cinemas')),
(4, 80, 8, (SELECT theater_id FROM theaters WHERE "name" = 'MegaPlex Cinemas')),
(5, 70, 7, (SELECT theater_id FROM theaters WHERE "name" = 'MegaPlex Cinemas'));

INSERT INTO genres ("name") VALUES
('Action'),
('Adventure'),
('Comedy'),
('Drama'),
('Fantasy'),
('Horror'),
('Mystery'),
('Romance'),
('Sci-Fi'),
('Thriller'),
('Animation'),
('Documentary'),
('Family'),
('Crime'),
('Musical'),
('Biography');

SELECT t.*
FROM public.seats t;

-- Insert seat for every room
INSERT INTO seats (seat_number, "row", availability, room_id)
SELECT
    ((room.seats / room.rows) * (row_index - 1)) + seat_in_row AS seat_number,
    row_index AS row_number,
    true AS availability, -- Assuming all seat are initially available
    room.room_id
FROM
    rooms AS room
        CROSS JOIN LATERAL
        generate_series(1, room.rows) AS row_index
        CROSS JOIN LATERAL
        generate_series(1, room.seats / room.rows) AS seat_in_row;

