-- Ensure the pgcrypto extension is enabled for UUID functions, if needed
-- CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO auths (email, hashed_password, role, profile_readable_id)
VALUES 
    ('donor1@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '9c65559f-5837-43ba-8766-5f4015452dc8'),
    ('donor2@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '3787d0a4-8950-462f-bcc6-ce7a2bc953a4'),
    ('donor3@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '7f4cccae-c094-435d-b0a7-c6e30e81923e'),
    ('donor4@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '14219c04-f57c-4f48-86c7-e4e3704b9147'),
    ('donor5@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '2e13bf2e-bfb9-4083-848e-d0a5b84560e1'),

    ('donor6@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '1241d452-3286-4468-90e8-3a5d89c07d88'),
    ('donor7@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '0256d979-6a16-4fbd-b07b-14f2351ea9e0'),
    ('donor8@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '5fd08b80-117c-45bc-8605-41dbe91136c7'),
    ('donor9@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '02d03f57-34b0-416f-8f30-e41d8941198e'),
    ('donor10@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', 'e3805b3b-d81d-49e4-ba6e-bdf675c528e6'),

    ('donor11@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '50a15bd7-07fd-4bb5-ab67-e6f1ef068773'),
    ('donor12@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '418ab030-ae14-41e2-9d53-f0b05740d093'),
    ('donor13@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '29064973-3817-417c-a8aa-c215e7c2b053'),
    ('donor14@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', 'dcb95859-54b2-4b79-827d-e5edce001c3f'),
    ('donor15@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', 'f9c60aea-c911-493d-98af-bd2e636ace18'),

    ('donor16@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '114b72d5-25ed-40d2-89f7-0d14f223dd37'),
    ('donor17@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '616878f4-f8df-4ec7-b1fb-bd821bd36aec'),
    ('donor18@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', 'cb99bc0e-9418-425e-bb72-d844f9f5523b'),
    ('donor19@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '37ba7ed4-46d3-4ec8-94fc-f5a45e55c7f3'),
    ('donor20@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '68f28cb8-37e4-42b2-822b-fc017df4f640'),

    ('donor21@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', 'ecf874e7-be8d-4a18-8c7e-cf6cb56d6897'),
    ('donor22@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', 'eae027af-826a-48e3-bdc8-930ad7ab6aa8'),
    ('donor23@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', 'c29f60a6-0ad4-431c-8ed3-cde26371f049'),
    ('donor24@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', 'f60c6c16-1938-4016-a2b7-024a93a34927'),
    ('donor25@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', 'd3296b09-60e3-4346-9011-59c227e8670e'),

    ('donor26@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '4d3f1dea-a04b-4844-89a6-a6f4e525a7ab'),
    ('donor27@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', 'e75c9559-4f25-4aa2-b2fd-2e2a0ed1f220'),
    ('donor28@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '6410dd2d-9b86-4b8c-a44d-f32049ce2fef'),
    ('donor29@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', '66180a74-0bfb-4747-99dd-95d9cdaf8da1'),
    ('donor30@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'DONOR', 'd51aa6cb-a471-45f8-aa15-34924a3f5c95');


-- Insert charity organization accounts into the auths table
INSERT INTO auths (readable_id, email, hashed_password, role, profile_readable_id)
VALUES 
    ('cf81bbc7-8990-4b9c-847d-ea0f3bc3e28a', 'charity1@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'CHARITY', 'a364e55f-2b2d-4180-b29c-c1385cb3c27e'),
    ('5b08c663-a605-43e1-86dd-0941021f6383', 'charity2@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'CHARITY', '209eddd5-561f-4605-9103-07ed9f8e36b2'),
    ('921a99c0-47cd-4507-bc7c-bc284d15c7b9', 'charity3@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'CHARITY', '713a7a9f-457f-4159-bb9e-b9835f9313f3'),
    ('47fe1f6d-05c4-4ff4-9e5d-91f47b585380', 'charity4@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'CHARITY', '6c648bad-4036-4240-9685-199eb7c4e164'),
    ('f61746df-3de1-47bd-88ee-3705df97dbac', 'charity5@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'CHARITY', '87e36b91-1047-4b3a-ab30-617b03df537a'),
    ('c0132204-680d-450b-9fce-6f4f91f90dcf', 'charity6@example.com', '$2a$10$AdA8zvXzhbzs1eodny4NkuR2DkUWN27j7WU5n/z0S7LTvlmYlsJ8y', 'CHARITY', '3c1e982f-bb78-42c9-8829-690f5ee842ed');


