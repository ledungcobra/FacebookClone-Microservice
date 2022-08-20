import React from 'react';

function DateOfBirthSelect({bDay, bMonth, bYear, handleRegisterChange}) {
    const currentYear = new Date().getFullYear();
    function getDayInMonth(month, year) {
        return new Date(year, month, 0).getDate();
    }
    return (
        <div className="reg_grid">
            <select name="bDay" value={bDay}>
                {
                    Array(getDayInMonth(bMonth, bYear)).fill(0).map((_, i) => <option
                        key={i + 1}
                        value={i + 1 + ''}>{i + 1}</option>)
                }
            </select>
            <select name="bMonth" value={bMonth}>
                {Array(12).fill(0).map((_, i) => <option key={i + 1}
                                                         value={i + 1 + ''}>{i + 1}</option>)}
            </select>
            <select name="bYear" value={bYear}>
                {Array(currentYear - 1900).fill(0).map((_, i) => <option key={i + 1900 + 1}
                                                                         value={i + 1900 + 1 + ''}>{i + 1900 + 1}</option>)}
            </select>
        </div>
    );
}

export default DateOfBirthSelect;