INSERT INTO organization (id, name, description, type) VALUES
    ('6d0f934a-05e7-4324-a001-7e71399494d7', 'org name 1', 'org desc 1', 'IE'),
    ('90d6655b-20f2-4b23-80ab-971e0be6d1cc', 'org name 2', 'org desc 2', 'LLC'),
    ('6cbebfeb-4fbd-4b0b-af93-af385cb5cce0', 'org name 3', 'org desc 3', 'JSC');

INSERT INTO employee (id, username, first_name, last_name) VALUES
    ('5f44af3a-ef52-4e50-8176-992422085100', 'user1', 'first name 1', 'last name 1'),
    ('839b89b8-0645-4827-bc42-893bdd94a67a', 'user2', 'first name 2', 'last name 2'),
    ('e7d29b7e-2c35-4c97-80d8-085f9f1caea4', 'user3', 'first name 3', 'last name 3'),
    ('5a3658e6-f31c-4271-b6a4-c7435b66f816', 'user4', 'first name 4', 'last name 4'),
    ('5128c85f-4d37-4f6b-a741-0479a66bd245', 'user5', 'first name 5', 'last name 5'),
    ('5690cc38-0ab3-4d12-8147-e141d7f6f8a1', 'user6', 'first name 6', 'last name 6');

INSERT INTO organization_responsible (id, organization_id, user_id) VALUES
    ('3b1abea3-8c46-4056-a9ff-041a656b5dd5', '6d0f934a-05e7-4324-a001-7e71399494d7', '5f44af3a-ef52-4e50-8176-992422085100'),
    ('27853b4f-eb67-4146-b130-500ed82f1be4', '90d6655b-20f2-4b23-80ab-971e0be6d1cc', '839b89b8-0645-4827-bc42-893bdd94a67a'),
    ('0570cc3f-2e78-4fee-9d74-a28437d5f087', '6cbebfeb-4fbd-4b0b-af93-af385cb5cce0', 'e7d29b7e-2c35-4c97-80d8-085f9f1caea4'),
    ('4c855981-bcbb-424e-939d-84dc4ee61582', '6cbebfeb-4fbd-4b0b-af93-af385cb5cce0', '5a3658e6-f31c-4271-b6a4-c7435b66f816');