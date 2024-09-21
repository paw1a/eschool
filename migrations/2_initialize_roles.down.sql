revoke all on all tables in schema public from root;
revoke all on all tables in schema public from guest;
revoke all on all tables in schema public from authenticated;
drop role root, guest, authenticated;
