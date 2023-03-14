package services

import "github.com/otisnado/fofo-server/models"

func GetPolicies() ([]models.Policy, error) {
	var policies []models.Policy
	if err := models.DB.Find(&policies).Error; err != nil {
		return nil, err
	}
	return policies, nil
}

func GetPolicyById(policyId uint) (*models.Policy, error) {
	var policy *models.Policy
	if err := models.DB.Where("id = ?", policyId).First(&policy).Error; err != nil {
		return nil, err
	}

	return policy, nil
}

func CreatePolicy(policy *models.Policy) (bool, error) {
	if err := models.DB.Create(&policy).Error; err != nil {
		return false, err
	}

	return true, nil
}

func UpdatePolicy(policyId uint, policyInput models.PolicyUpdate) (*models.PolicyUpdate, error) {
	var policy models.Policy

	if err := models.DB.Where("id = ?", policyId).Model(&policy).Updates(policyInput).Error; err != nil {
		return nil, err
	}

	return &policyInput, nil
}

func DeletePolicy(policyId uint) (bool, error) {
	var policy models.Policy
	if err := models.DB.Where("id = ?", policyId).First(&policy).Error; err != nil {
		return false, err
	}
	if err := models.DB.Delete(&policy).Error; err != nil {
		return false, err
	}

	return true, nil
}

func GetPoliciesByRoleId(roleId uint) ([]models.Policy, error) {
	var policies []models.Policy
	if err := models.DB.Where("authorized_role = ?", roleId).Find(&policies).Error; err != nil {
		return nil, err
	}

	return policies, nil
}
