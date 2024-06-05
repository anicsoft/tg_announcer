import { ActionIcon, BoxComponentProps, Button, CloseButton, Divider, Flex, Group, Input, MantineStyleProp, Select, Stack, Switch, Text, TextInput, rem } from '@mantine/core'
import { IconAt, IconClock } from '@tabler/icons-react'
import React, { useEffect, useRef, useState } from 'react'
import RichTexInput from '../ui/RichTexInput';
import { useForm } from '@mantine/form';
import { useQuery } from '@tanstack/react-query';
import { DatePickerInput, DateTimePicker, TimeInput } from '@mantine/dates';

export default function AdminConsole() {
  const [hasPromocode, setHasPromocode] = useState(false)
  const [singleDayOffer, setSingleDayOffer] = useState(true)
  const [dateRange, setDateRange] = useState([])
  const [date, setDate] = useState()
  const [innerDate, setInnerDate] = useState(Date.now())
  const [startTime, setStartTime] = useState()
  const [endTime, setEndTime] = useState()
  const startTimeRef = useRef<HTMLInputElement>(null);
  const endTimeRef = useRef<HTMLInputElement>(null);


  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      title: '',
      hasPromocode: hasPromocode,
      promocode: '',
      content: '',
      dateRange: [],
      date: null,
      startTime: '',
      endTime: '',
      offerCategory: '',
      singleDayOffer: true, // DK if it is needed in the form
    },

    validate: {
      // email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
    },
  });


  const handlePromocodeSwitchChange = (isChecked) => {
    form.setValues({
      hasPromocode: isChecked
    })
    setHasPromocode(isChecked)
  }

  const handleSingleDayOfferSwitchChange = (isChecked) => {
    form.setValues({
      singleDayOffer: isChecked
    })
    setSingleDayOffer(isChecked)
  }

  const handleDateRangeChange = (value) => {
    console.log(value);

    form.setValues({
      dateRange: value
    })
    setDateRange(value)
  }

  const handleDateChange = (value) => {
    console.log(value);

    form.setValues({
      date: value
    })

    setDateRange([new Date(value), new Date(value)])
    setDate(value)
  }

  const handleTimeChange = (value: string, start: boolean) => {
    console.log("handle time change");
    console.log(value);

    const timeParts = value.split(':');
    const hours = parseInt(timeParts[0], 10);
    const minutes = parseInt(timeParts[1], 10);

    // Create a new Date object with the existing date
    // const newDate = new Date(innerDate)
    // console.log(date);

    const existingDate = dateRange;
    if (start) {
      setStartTime(value)
      form.setValues({
        startTime: value
      })

      const newDate = new Date(existingDate[0]);

      // Set the time components
      newDate.setHours(hours);
      newDate.setMinutes(minutes);

      const newDateRange = [...dateRange]
      newDateRange.splice(0, 1, newDate)
      form.setValues({
        dateRange: newDateRange
      })
      setDateRange(newDateRange)
    } else {

      setEndTime(value)

      form.setValues({
        endTime: value
      })


      const newDate = new Date(existingDate[1]);
      // Set the time components
      newDate.setHours(hours);
      newDate.setMinutes(minutes);

      const newDateRange = [...dateRange]
      newDateRange.splice(1, 1, newDate)
      form.setValues({
        dateRange: newDateRange
      })
      setDateRange(newDateRange)

    }
  }

  // useEffect(() => {
  //   console.log({ ...form.getValues() }.hasPromocode);
  //   console.log({ ...form.getInputProps('hasPromocode') });

  //   setHasPromocode({ ...form.getValues() }.hasPromocode)

  //   // return () => {
  //   //   second
  //   // }
  // }, [hasPromocode])

  const { data: offerCategories } = useQuery({
    queryKey: ['offerCategories'],
    queryFn: () => fetch('http://localhost:8888/categories/offer').then((res) =>
      res.json(),
    )
  })

  const startTimePickerControl = (
    <ActionIcon variant="subtle" color="gray" onClick={() => startTimeRef.current?.showPicker()}>
      <IconClock style={{ width: rem(16), height: rem(16) }} stroke={1.5} />
    </ActionIcon>
  );
  const endTimePickerControl = (
    <ActionIcon variant="subtle" color="gray" onClick={() => endTimeRef.current?.showPicker()}>
      <IconClock style={{ width: rem(16), height: rem(16) }} stroke={1.5} />
    </ActionIcon>
  );


  const DependentFields1 = React.memo(() => {
    return (
      <>
        {hasPromocode && <TextInput
          // required
          // mt="xl"
          style={inputLabelStyles}
          label="Promocode"
          // disabled={!{ ...form.getValues() }.hasPromocode}
          key={form.key('promocode')}
          {...form.getInputProps('promocode')}
        // placeholder="Custom layout"
        // description="Write offer content here"
        // error="both below the input"
        // inputWrapperOrder={['label', 'input', 'description', 'error']}
        />}
      </>
    );
  });

  const DependentFields2 = React.memo(() => {
    return (
      <>
        {/* {!singleDayOffer && <Group
          flex={1}
        >

          <TimeInput
            label="Start time"
            withSeconds={false}
            flex={1}
            // description="Input description"
            placeholder="Start time"
            ref={startTimeRef}
            rightSection={startTimePickerControl}
            value={startTime}
            onChange={(value) => handleTimeChange(value, true)}
          />
          <TimeInput
            label="End time"
            withSeconds={false}
            flex={1}
            // description="Input description"
            placeholder="End time"
            ref={endTimeRef}
            rightSection={endTimePickerControl}
            value={endTime}
            onChange={(value) => handleTimeChange(value, false)}
          />
        </Group>} */}
        <Input.Wrapper label="Offer content">
          {/* <Input placeholder="Input inside Input.Wrapper" /> */}
          <RichTexInput content={{ ...form.getInputProps('content') }.defaultValue} onChange={handleOfferContentChange}></RichTexInput>
        </Input.Wrapper >
      </>
    );
  });

  const getTimeSlots = (startTime = "00:00", endTime = "24:00") => {
    const start = 0; // Start time in hours
    const end = 24; // End time in hours

    const slots = Array.from({ length: (end - start) * 2 }, (_, index) => {
      const hour = start + Math.floor(index / 2);
      const minute = index % 2 === 0 ? '00' : '30';
      return `${String(hour).padStart(2, '0')}:${minute}`;
    });

    const filteredSlots = slots.filter(slot => {
      return (startTime === "00:00" || slot > startTime) && (endTime === "24:00" || slot < endTime);
    });
    console.log(filteredSlots);

    return filteredSlots
  }

  const StartTimeList = () => {
    const slots = getTimeSlots(undefined, endTime)

    return (
      <>
        <datalist id="startTime">
          {slots.map((slot) =>
            <option key={slot} value={slot} >{slot}</option>
          )}
        </datalist>
      </>
    )
  }
  const EndTimeList = () => {
    const slots = getTimeSlots(startTime, undefined)
    return (
      <>
        <datalist id="endTime">
          {slots.map((slot) =>
            <option key={slot} value={slot} >{slot}</option>
          )}
          {/* <option value="10:00">10:00</option>
        <option value="11:00">11:00</option>
        <option value="12:00">12:00</option> */}
        </datalist>
      </>
    )
  }

  const handleOfferContentChange = (value) => {
    form.setValues({
      content: value
    })
  }


  const inputLabelStyles: MantineStyleProp = { textAlign: "left" }
  return (
    <>
      <Text size={"lg"}>Add offer</Text>
      <form onSubmit={form.onSubmit((values) => { console.log(values), console.log(hasPromocode) })}>

        <Flex
          align={"flex-start"}
          direction={"column"}
          px="lg"
          gap={48}
        >
          <Stack
            bg="var(--mantine-color-body)"
            align="stretch"
            justify="center"
            gap="md"
          >
            <TextInput
              required
              label="Title"
              style={inputLabelStyles}
              key={form.key('title')}
              {...form.getInputProps('title')}
            // placeholder="Title"
            // description="Description below the input"
            // inputWrapperOrder={['label', 'error', 'input', 'description']}
            />
            <Select
              label="Offer type"
              placeholder="Pick value"
              required
              key={form.key('offerCategory')}
              data={offerCategories ? offerCategories.map(cat => cat.name) : []}
              {...form.getInputProps('offerCategory')}
            />
            <Switch
              // defaultChecked
              style={inputLabelStyles}
              label="Offer has a promocode"
              key={form.key('hasPromocode')}


              type='checkbox'
              checked={hasPromocode}
              // value={hasPromocode}
              onChange={(event) => handlePromocodeSwitchChange(event.currentTarget.checked)}
            // {...form.getInputProps(`hasPromocode`, { type: 'checkbox' })}
            />
            <DependentFields1></DependentFields1>

          </Stack>
          <Stack
            bg="var(--mantine-color-body)"
            align="stretch"
            justify="center"
            gap="md"
          >
            <Input.Wrapper
              required

              style={inputLabelStyles}
              // label="Some label"
              // description="Some description"
              label="Offer content"
              // placeholder="Custom layout"
              description="Write offer content here"

              key={form.key('content')}
              {...form.getInputProps('content')}
            // error="Some error"
            >
              {/* <RichTexInput content={form.getValues().content} onChange={handleOfferContentChange}></RichTexInput> */}
              <RichTexInput initcontent={form.getValues().content}></RichTexInput>
              {/* <DependentFields2></DependentFields2> */}
              {/* <Input<any> component={TextEditor} /> */}
            </Input.Wrapper>
          </Stack>
          <Divider my="md" color='#000 ' h={1} w={"100%"} />
          <Stack
            bg="var(--mantine-color-body)"
            align="stretch"
            justify="center"
            gap="md"
            w={"100%"}
          >
            <Text size={"lg"}>Offer duration</Text>
            <DatePickerInput
              label="Pick date"
              placeholder="Pick single date"
              value={date}
              onChange={(value) => handleDateChange(value)}
            />
            <Switch
              style={inputLabelStyles}
              label="Whole day"
              key={form.key('singleDayOffer')}
              defaultChecked={singleDayOffer}
              type='checkbox'
              checked={singleDayOffer}
              // value={hasPromocode}
              onChange={(event) => handleSingleDayOfferSwitchChange(event.currentTarget.checked)}
            />
            {!singleDayOffer && <Group
              flex={1}
            >

              <TimeInput
                label="Start time"
                withSeconds={false}
                flex={1}
                list="startTime"
                disabled={!form.getValues().date}
                maxTime={endTime}
                // description="Input description"
                placeholder="Start time"
                ref={startTimeRef}
                rightSection={startTimePickerControl}
                value={form.getValues().startTime}
                onChange={(event) => handleTimeChange(event.currentTarget.value, true)}
              // {...form.getInputProps('startTime')}
              />
              <TimeInput
                label="End time"
                list="endTime"
                withSeconds={false}
                // maxTime={"23:59"}
                minTime={startTime}
                disabled={!form.getValues().date}
                flex={1}
                // description="Input description"
                placeholder="End time"
                ref={endTimeRef}
                rightSection={endTimePickerControl}
                // value={endTime}
                value={form.getValues().endTime}
                onChange={(event) => handleTimeChange(event.currentTarget.value, false)}
              />
              {/* {timeList} */}
              <StartTimeList></StartTimeList>
              <EndTimeList></EndTimeList>
            </Group>}
            {/* <DependentFields2></DependentFields2> */}

          </Stack>
          <Button type="submit">Submit</Button>
        </Flex>
      </form>

    </>
  )
}
