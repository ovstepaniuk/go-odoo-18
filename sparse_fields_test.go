package odoo

// SparseFieldsTest represents sparse_fields.test model.
type SparseFieldsTest struct {
	Boolean     *Bool       `xmlrpc:"boolean,omitempty"`
	Char        *String     `xmlrpc:"char,omitempty"`
	CreateDate  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One   `xmlrpc:"create_uid,omitempty"`
	Data        interface{} `xmlrpc:"data,omitempty"`
	DisplayName *String     `xmlrpc:"display_name,omitempty"`
	Float       *Float      `xmlrpc:"float,omitempty"`
	Id          *Int        `xmlrpc:"id,omitempty"`
	Integer     *Int        `xmlrpc:"integer,omitempty"`
	Partner     *Many2One   `xmlrpc:"partner,omitempty"`
	Selection   *Selection  `xmlrpc:"selection,omitempty"`
	WriteDate   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// SparseFieldsTests represents array of sparse_fields.test model.
type SparseFieldsTests []SparseFieldsTest

// SparseFieldsTestModel is the odoo model name.
const SparseFieldsTestModel = "sparse_fields.test"

// Many2One convert SparseFieldsTest to *Many2One.
func (st *SparseFieldsTest) Many2One() *Many2One {
	return NewMany2One(st.Id.Get(), "")
}

// CreateSparseFieldsTest creates a new sparse_fields.test model and returns its id.
func (c *Client) CreateSparseFieldsTest(st *SparseFieldsTest) (int64, error) {
	ids, err := c.CreateSparseFieldsTests([]*SparseFieldsTest{st})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSparseFieldsTest creates a new sparse_fields.test model and returns its id.
func (c *Client) CreateSparseFieldsTests(sts []*SparseFieldsTest) ([]int64, error) {
	var vv []interface{}
	for _, v := range sts {
		vv = append(vv, v)
	}
	return c.Create(SparseFieldsTestModel, vv, nil)
}

// UpdateSparseFieldsTest updates an existing sparse_fields.test record.
func (c *Client) UpdateSparseFieldsTest(st *SparseFieldsTest) error {
	return c.UpdateSparseFieldsTests([]int64{st.Id.Get()}, st)
}

// UpdateSparseFieldsTests updates existing sparse_fields.test records.
// All records (represented by ids) will be updated by st values.
func (c *Client) UpdateSparseFieldsTests(ids []int64, st *SparseFieldsTest) error {
	return c.Update(SparseFieldsTestModel, ids, st, nil)
}

// DeleteSparseFieldsTest deletes an existing sparse_fields.test record.
func (c *Client) DeleteSparseFieldsTest(id int64) error {
	return c.DeleteSparseFieldsTests([]int64{id})
}

// DeleteSparseFieldsTests deletes existing sparse_fields.test records.
func (c *Client) DeleteSparseFieldsTests(ids []int64) error {
	return c.Delete(SparseFieldsTestModel, ids)
}

// GetSparseFieldsTest gets sparse_fields.test existing record.
func (c *Client) GetSparseFieldsTest(id int64) (*SparseFieldsTest, error) {
	sts, err := c.GetSparseFieldsTests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sts)[0]), nil
}

// GetSparseFieldsTests gets sparse_fields.test existing records.
func (c *Client) GetSparseFieldsTests(ids []int64) (*SparseFieldsTests, error) {
	sts := &SparseFieldsTests{}
	if err := c.Read(SparseFieldsTestModel, ids, nil, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSparseFieldsTest finds sparse_fields.test record by querying it with criteria.
func (c *Client) FindSparseFieldsTest(criteria *Criteria) (*SparseFieldsTest, error) {
	sts := &SparseFieldsTests{}
	if err := c.SearchRead(SparseFieldsTestModel, criteria, NewOptions().Limit(1), sts); err != nil {
		return nil, err
	}
	return &((*sts)[0]), nil
}

// FindSparseFieldsTests finds sparse_fields.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSparseFieldsTests(criteria *Criteria, options *Options) (*SparseFieldsTests, error) {
	sts := &SparseFieldsTests{}
	if err := c.SearchRead(SparseFieldsTestModel, criteria, options, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSparseFieldsTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSparseFieldsTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SparseFieldsTestModel, criteria, options)
}

// FindSparseFieldsTestId finds record id by querying it with criteria.
func (c *Client) FindSparseFieldsTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SparseFieldsTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
