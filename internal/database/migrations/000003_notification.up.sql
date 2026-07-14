CREATE TYPE status AS ENUM (
    'pending',
    'sent',
    'failed'
);

CREATE TABLE notifications(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    template_id UUID NOT NULL,
    data JSONB NOT NULL,
    status status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    FOREIGN KEY (user_id)
        REFRENCES users(id)
        ON DELETE CASCADE

    FOREIGN KEY (template_id)
        REFRENCES templates(id)
        ON DELETE CASCADE
);