CREATE INDEX drivers_status_idx ON drivers(status);
CREATE INDEX vehicles_driver_active_idx ON vehicles(driver_id, active);
CREATE INDEX compliance_driver_time_idx ON compliance_checks(driver_id, checked_at DESC);
CREATE INDEX audit_entity_idx ON audit_logs(entity_type, entity_id, created_at DESC);
