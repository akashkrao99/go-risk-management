package risks

func isValidStatus(status RiskStatus) bool {
    for _, item := range allowedRiskStatuses {
        if item == status {
            return true
        }
    }
    return false
}