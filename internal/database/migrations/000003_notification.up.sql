CREATE TYPE status AS ENUM (
    'pending',
    'sent',
    'failed'
);

CREATE TABLE notifications(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL
    template_id UUID NOT NULL
    status status NOT NULL DEFAULT 'pending'

    FOREIGN KEY (user_id)
        REFRENCES users(id)
        ON DELETE CASCADE

    FOREIGN KEY (template_id)
        REFRENCES templates(id)
        ON DELETE CASCADE
);