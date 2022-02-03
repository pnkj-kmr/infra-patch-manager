package entity

// Conf holds the all configuration variables of application
// Variables loaded from env file and os env by viber
type _conf struct {
	p string `mapstructure:"DIR_PATCH"`
	r string `mapstructure:"DIR_PATCH_ROLLBACK"`
	a string `mapstructure:"DIR_ASSETS"`
}

func (c *_conf) AssetPath() string    { return c.a }
func (c *_conf) PatchPath() string    { return c.p }
func (c *_conf) RollbackPath() string { return c.r }
