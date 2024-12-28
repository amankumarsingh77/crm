
CREATE TYPE user_role AS ENUM ('admin', 'user');
CREATE TYPE application_stage AS ENUM ('counselling', 'college_selection', 'application_status', 'visa', 'loan', 'complete');
CREATE TYPE status_type AS ENUM ('pending', 'in_progress', 'completed');

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       first_name VARCHAR(100) NOT NULL,
                       last_name VARCHAR(100) NOT NULL,
                       phone VARCHAR(20) UNIQUE NOT NULL,
                       role user_role NOT NULL DEFAULT 'user',
                       is_active BOOLEAN NOT NULL DEFAULT true,
                       last_login TIMESTAMP WITH TIME ZONE,
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE applications (
                              id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                              user_id UUID NOT NULL REFERENCES users(id),
                              current_stage application_stage NOT NULL DEFAULT 'counselling',
                              status status_type NOT NULL DEFAULT 'pending',
                              college_name VARCHAR(255),
                              course_name VARCHAR(255),
                              intake_date DATE,
                              created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE stage_progress (
                                id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                application_id UUID NOT NULL REFERENCES applications(id),
                                stage application_stage NOT NULL,
                                status status_type NOT NULL DEFAULT 'pending',
                                start_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                completion_date TIMESTAMP WITH TIME ZONE,
                                created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                CONSTRAINT fk_application FOREIGN KEY (application_id) REFERENCES applications(id) ON DELETE CASCADE
);

CREATE TABLE documents (
                           id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                           stage_progress_id UUID NOT NULL REFERENCES stage_progress(id),
                           document_type VARCHAR(50) NOT NULL,
                           file_name VARCHAR(255) NOT NULL,
                           s3_path VARCHAR(512) NOT NULL,
                           file_size INTEGER NOT NULL,
                           content_type VARCHAR(100) NOT NULL,
                           uploaded_by UUID NOT NULL REFERENCES users(id),
                           is_verified BOOLEAN NOT NULL DEFAULT false,
                           verified_by UUID REFERENCES users(id),
                           verified_at TIMESTAMP WITH TIME ZONE,
                           created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           CONSTRAINT fk_stage_progress FOREIGN KEY (stage_progress_id) REFERENCES stage_progress(id) ON DELETE CASCADE,
                           CONSTRAINT fk_uploaded_by FOREIGN KEY (uploaded_by) REFERENCES users(id),
                           CONSTRAINT fk_verified_by FOREIGN KEY (verified_by) REFERENCES users(id)
);

CREATE TABLE stage_notes (
                             id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                             stage_progress_id UUID NOT NULL REFERENCES stage_progress(id),
                             note TEXT NOT NULL,
                             created_by UUID NOT NULL REFERENCES users(id),
                             created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             CONSTRAINT fk_stage_progress FOREIGN KEY (stage_progress_id) REFERENCES stage_progress(id) ON DELETE CASCADE,
                             CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id)
);

CREATE TABLE notifications (
                               id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                               user_id UUID NOT NULL REFERENCES users(id),
                               application_id UUID NOT NULL REFERENCES applications(id),
                               title VARCHAR(255) NOT NULL,
                               message TEXT NOT NULL,
                               is_read BOOLEAN NOT NULL DEFAULT false,
                               read_at TIMESTAMP WITH TIME ZONE,
                               created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                               CONSTRAINT fk_application FOREIGN KEY (application_id) REFERENCES applications(id) ON DELETE CASCADE
);


CREATE TABLE audit_logs (
                            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                            user_id UUID NOT NULL REFERENCES users(id),
                            action VARCHAR(100) NOT NULL,
                            entity_type VARCHAR(50) NOT NULL,
                            entity_id UUID NOT NULL,
                            old_value JSONB,
                            new_value JSONB,
                            ip_address VARCHAR(45) NOT NULL,
                            created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);
CREATE INDEX idx_applications_user_id ON applications(user_id);
CREATE INDEX idx_applications_status ON applications(status);
CREATE INDEX idx_stage_progress_application_id ON stage_progress(application_id);
CREATE INDEX idx_stage_progress_status ON stage_progress(status);
CREATE INDEX idx_documents_stage_progress_id ON documents(stage_progress_id);
CREATE INDEX idx_notifications_user_id ON notifications(user_id);
CREATE INDEX idx_notifications_is_read ON notifications(is_read);
CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_entity_id ON audit_logs(entity_id);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_applications_updated_at
    BEFORE UPDATE ON applications
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_stage_progress_updated_at
    BEFORE UPDATE ON stage_progress
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_documents_updated_at
    BEFORE UPDATE ON documents
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_stage_notes_updated_at
    BEFORE UPDATE ON stage_notes
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();